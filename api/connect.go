package api

import (
	"encoding/json"
	"github.com/stefanopulze/daitem/data"
	"log"
)

type connectRequest struct {
	ConnectionType string `json:"connectionType"`
	MasterCode     string `json:"masterCode"`
	SessionId      string `json:"sessionId"`
	TransmitterId  string `json:"transmitterId"`
}

type connectResponse struct {
	Status

	ClientIpAddress string `json:"clientIpAddress"`
	FirmwareVersion string `json:"firmwareVersion"`
	GprsConnection  string `json:"gprsConnection"`
	Groups          []int  `json:"groups"`
	SystemIpAddress string `json:"systemIpAddress"`
	SystemState     string `json:"systemState"`
	TTMSessionId    string `json:"ttmSessionId"`
}

func (api *Api) Connect() (*data.DeviceInfo, error) {
	request := connectRequest{
		ConnectionType: api.context.ConnectionType,
		MasterCode:     api.context.MasterCode,
		SessionId:      api.context.SessionId,
		TransmitterId:  api.context.TransmitterId,
	}

	response, err := api.sendPost("/authenticate/connect", request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var respJson connectResponse
	if err := json.NewDecoder(response.Body).Decode(&respJson); err != nil {
		return nil, err
	}

	if err := checkValidStatus(respJson.Status); err != nil {
		return nil, err
	}

	api.context.TTMSessionId = respJson.TTMSessionId
	log.Printf("Found new ttmSessionId: %s", respJson.TTMSessionId)

	return &data.DeviceInfo{
		ClientIpAddress: respJson.ClientIpAddress,
		FirmwareVersion: respJson.FirmwareVersion,
		GprsConnection:  respJson.GprsConnection,
		Groups:          respJson.Groups,
		SystemIpAddress: respJson.SystemIpAddress,
		SystemState:     respJson.SystemState,
		TTMSessionId:    respJson.TTMSessionId,
	}, nil
}
