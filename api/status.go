package api

import (
	"encoding/json"
	"errors"
	"github.com/stefanopulze/daitem/data"
)

type statusRequest struct {
	CentralId    string `json:"centralId,omitempty"`
	TtmSessionId string `json:"ttmSessionId"`
}

type statusResponse struct {
	Status

	CommandStatus string `json:"commandStatus"`
	Defaults      string `json:"defaults"`
	Group         []int  `json:"group"`
	SystemState   string `json:"systemState"`
}

func (api *Api) SystemState() (*data.DeviceStatus, error) {
	if api.context.TTMSessionId == "" {
		errors.New("invalid ttmSessionId. Please connect first")
	}

	request := statusRequest{
		TtmSessionId: api.context.TTMSessionId,
	}

	response, err := api.sendPost("/status/getSystemState", request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var status statusResponse
	if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
		return nil, err
	}

	if err := checkValidStatus(status.Status); err != nil {
		return nil, err
	}

	println(status.CommandStatus)

	return &data.DeviceStatus{
		SystemState:   status.SystemState,
		CommandStatus: status.CommandStatus,
	}, nil
}

func (api *Api) CurrentState() (*data.DeviceStatus, error) {
	if api.context.TTMSessionId == "" {
		return nil, errors.New("invalid ttmSessionId. Please connect first")
	}

	request := statusRequest{
		CentralId:    api.context.CentralId,
		TtmSessionId: api.context.TTMSessionId,
	}

	response, err := api.sendPost("/status/getCurrentState", request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var status statusResponse
	if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
		return nil, err
	}

	if err := checkValidStatus(status.Status); err != nil {
		return nil, err
	}

	return &data.DeviceStatus{
		SystemState:   status.SystemState,
		CommandStatus: status.CommandStatus,
	}, nil
}
