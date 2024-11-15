package api

import (
	"encoding/json"
	"github.com/stefanopulze/daitem/data"
)

type configurationRequest struct {
	Username  string `json:"username"`
	SessionId string `json:"sessionId"`
}

type configurationResponse struct {
	Status

	TransmitterId    string `json:"transmitterId"`
	CentralId        string `json:"centralId"`
	SMSEnabled       bool   `json:"smsEnabled"`
	AlarmPhoneNumber string `json:"alarmPhoneNumber"`
	ConnectionType   string `json:"connectionType"`
	MonitorType      string `json:"monitoringType"`
}

func (api *Api) Configuration() (*data.DeviceConfiguration, error) {
	if api.context.TransmitterId != "" && api.context.CentralId != "" && api.context.ConnectionType != "" {
		return &data.DeviceConfiguration{
			TransmitterId:  api.context.TransmitterId,
			CentralId:      api.context.CentralId,
			ConnectionType: api.context.ConnectionType,
			UseCache:       true,
		}, nil
	}

	requestData := configurationRequest{
		Username:  api.context.Username,
		SessionId: api.context.SessionId,
	}

	response, err := api.sendPost("/configuration/getConfiguration", requestData)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	var configuration configurationResponse

	if err := json.NewDecoder(response.Body).Decode(&configuration); err != nil {
		return nil, err
	}

	if err := checkValidStatus(configuration.Status); err != nil {
		return nil, err
	}

	api.context.TransmitterId = configuration.TransmitterId
	api.context.CentralId = configuration.CentralId
	api.context.ConnectionType = configuration.ConnectionType

	return &data.DeviceConfiguration{
		TransmitterId:  api.context.TransmitterId,
		CentralId:      api.context.CentralId,
		ConnectionType: api.context.ConnectionType,
		UseCache:       false,
	}, nil
}
