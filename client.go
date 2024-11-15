package daitem

import (
	"github.com/stefanopulze/daitem/api"
	"github.com/stefanopulze/daitem/context"
	"github.com/stefanopulze/daitem/data"
	"github.com/stefanopulze/daitem/storage"
	"log"
	"time"
)

type Client struct {
	Api        *api.Api
	storage    storage.Storage
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

	if savedContext, err := context.Load(*options.Storage); err == nil {
		ctx.Merge(savedContext)
	}

	return &Client{
		Api:     api.NewApi(&ctx),
		storage: *options.Storage,
		context: &ctx,
	}
}

func (client *Client) Status() (bool, error) {
	if err := client.ensureSession(); err != nil {
		return false, err
	}

	status, err := client.Api.CurrentState()
	if err != nil {
		return false, err
	}

	return status.SystemState == "on", nil
}

func (client *Client) TurnAlarm(on bool) error {
	if err := client.ensureSession(); err != nil {
		return err
	}

	if _, err := client.Api.TurnAlarm(on); err != nil {
		return err
	}

	log.Printf("Turning alarm: %t", on)
	return nil
}

func (client *Client) Info() (*data.DeviceInfo, error) {
	if err := client.ensureSession(); err != nil {
		return nil, err
	}

	return client.deviceInfo, nil
}

func (client *Client) ensureSession() error {
	if client.context.IsExpired() {
		// Get new sessionId
		session, err := client.Api.Login()
		if err != nil {
			return err
		} else if !session.UseCache {
			client.storage.Write(context.SessionId, []byte(session.SessionId))
			client.storage.Write(context.SessionTime, []byte(session.SessionTime.Format(time.RFC3339)))
		}

		log.Printf("Session id: %s", session.SessionId)

		// Get CentralId, TrasmitterId and ConnectionType if needed
		configuration, err := client.Api.Configuration()
		if err != nil {
			return err
		} else if !configuration.UseCache {
			client.storage.Write(context.CentralId, []byte(configuration.CentralId))
			client.storage.Write(context.TransmitterId, []byte(configuration.TransmitterId))
			client.storage.Write(context.ConnectionType, []byte(configuration.ConnectionType))
		}

		log.Printf("Device central id: %s", configuration.CentralId)
	}

	// Ger new ttmSessionId
	if e := client.Api.KeepAlive(); e != nil {
		info, err := client.Api.Connect()

		if err != nil {
			return err
		}

		client.storage.Write(context.TTMSessionId, []byte(info.TTMSessionId))

		client.deviceInfo = info
		log.Printf("Device info: %s", info.FirmwareVersion)
	}

	return nil
}
