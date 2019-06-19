package daitem

import (
	"github.com/stefanopulze/daitem/api"
	"github.com/stefanopulze/daitem/context"
	"github.com/stefanopulze/daitem/data"
	"github.com/stefanopulze/daitem/storage"
	"log"
)

type Client struct {
	api        *api.Api
	storage    *storage.Storage
	context    *context.Context
	deviceInfo *data.DeviceInfo
}

func NewClient(options *ClientOptions) *Client {
	ctx := context.Context{
		Username:      options.Username,
		Password:      options.Password,
		MasterCode:    options.MasterCode,
		CentralId:     options.CentralId,
		TransmitterId: options.TransmitterId,
	}

	return &Client{
		api:     api.NewApi(&ctx),
		storage: options.Storage,
	}
}

func (client *Client) Status() (bool, error) {
	if err := client.ensureSession(); err != nil {
		return false, err
	}

	return true, nil
}

func (client *Client) TurnAlarm(on bool) error {
	if err := client.ensureSession(); err != nil {
		return err
	}

	log.Print("Turning alarm on")
	return nil
}

func (client *Client) Info() (*data.DeviceInfo, error) {
	if err := client.ensureSession(); err != nil {
		return nil, err
	}

	return client.deviceInfo, nil
}

func (client *Client) ensureSession() error {
	// Get new sessionId
	session, err := client.api.Login()
	if err != nil {
		return err
	}

	log.Printf("Session id: %s", session.SessionId)

	// Get CentralId, TrasmitterId and ConnectionType if needed
	configuration, err := client.api.Configuration()
	if err != nil {
		return err
	}

	log.Printf("Device central id: %s", configuration.CentralId)

	// Ger new ttmSessionId
	if e := client.api.KeepAlive(); e != nil {
		info, err := client.api.Connect()

		if err != nil {
			return err
		}

		client.deviceInfo = info
		log.Printf("Device info: %s", info.FirmwareVersion)
	}

	return nil
}
