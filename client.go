package daitem

import (
	"errors"
	"fmt"
	"github.com/stefanopulze/daitem/api"
	"github.com/stefanopulze/daitem/http"
	"github.com/stefanopulze/daitem/session"
)

type Client struct {
	Api     *api.Client
	session *session.Session
}

func NewClient(options *Options) *Client {
	ses := session.New(options.Username, options.Password, options.MasterCode)
	appVersion := "5.0.0"
	if len(options.AppVersion) > 0 {
		appVersion = options.AppVersion
	}

	httpClient := http.NewClient("https://appv3.tt-monitor.com")
	httpClient.AddHeader("Accept", "application/json")
	httpClient.AddHeader("Content-Type", "application/json")
	httpClient.AddHeader("X-App-Name", "eNova")
	httpClient.AddHeader("X-Identity-Provider", "ATRALTECH_JANRAIN")
	httpClient.AddHeader("X-App-Platform", "ios")
	httpClient.AddHeader("X-App-Version", appVersion)
	httpClient.AddHeader("User-Agent", fmt.Sprintf("Daitem Secure/%s", appVersion))

	apiClient := api.NewClient(httpClient, ses)

	return &Client{
		Api:     apiClient,
		session: ses,
	}
}

// ListSystems List all user systems
func (c *Client) ListSystems() ([]api.System, error) {
	if err := c.ensureSession(); err != nil {
		return nil, err
	}

	return c.Api.Systems()
}

// GetSystemState get the state of system. Activated or not
func (c *Client) GetSystemState(systemId int) (*api.SystemState, error) {
	if err := c.ensureSession(); err != nil {
		return nil, err
	}

	return c.Api.SystemState(systemId)
}

// SystemConnect connect to transmitter and get a new ttmSessionId
func (c *Client) SystemConnect(systemId int) (*api.SystemConnect, error) {
	if err := c.ensureSession(); err != nil {
		return nil, err
	}

	return c.Api.Connect(systemId, c.session.GetMasterCode())
}

// SystemSendCommand send active or deactivate command to system.
// Function return error for connection error and for command nack
func (c *Client) SystemSendCommand(systemId int, activate bool) (*api.SystemState, error) {
	if err := c.ensureSession(); err != nil {
		return nil, err
	}

	response, err := c.Api.SystemSendCommand(systemId, activate)
	if err != nil {
		return nil, err
	}

	if response.CommandStatus != api.CommandOK {
		return nil, errors.New("system command error")
	}

	return response, err
}

// GetSystemConfiguration get system configurations like TransmitterId
func (c *Client) GetSystemConfiguration(systemId int) (*api.Configuration, error) {
	if err := c.ensureSession(); err != nil {
		return nil, err
	}

	return c.Api.GetSystemConfiguration(systemId)
}

// GetSystemInfo get information about central, transmitters, sirens, commands, firmwares ...
func (c *Client) GetSystemInfo(systemId int) (*api.SystemInfo, error) {
	if err := c.ensureSession(); err != nil {
		return nil, err
	}

	return c.Api.SystemRefresh(systemId)
}

// IsTransmitterConnect check if transmitter (or central) is connected
func (c *Client) IsTransmitterConnect(transmitterId string) (bool, error) {
	if err := c.ensureSession(); err != nil {
		return false, err
	}

	return c.Api.IsConnected(transmitterId)
}

// SystemUpdateAvailable check if there is some module can be updated on system
func (c *Client) SystemUpdateAvailable(systemId int) (bool, error) {
	if err := c.ensureSession(); err != nil {
		return false, err
	}

	response, err := c.Api.SystemRefresh(systemId)
	if err != nil {
		return false, err
	}

	return response.Central.Firmwares.UpdatableFirmwares(), nil
}

func (c *Client) ensureSession() error {
	if c.session.IsValid() {
		return nil
	}

	if len(c.session.GetRefreshToken()) > 0 {
		if data, err := c.Api.AuthWithRefreshToken(c.session.GetRefreshToken()); err == nil {
			c.session.Update(data.AccessToken, data.RefreshToken, data.ExpiresIn, data.UserId)
			return nil
		}
	}

	data, err := c.Api.AuthWithCredentials(c.session.GetUsername(), c.session.GetPassword())
	if err != nil {
		return err
	}

	c.session.Update(data.AccessToken, data.RefreshToken, data.ExpiresIn, data.UserId)
	return nil
}
