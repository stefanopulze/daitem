package api

import (
	"encoding/json"
	"errors"
	"io"
)

type commandRequest struct {
	CurrentGroup []string `json:"currentGroup"`
	Group        *string  `json:"group"`
	NbGroups     int      `json:"nbGroups"`
	SystemState  string   `json:"systemState"`
	TTMSessionId string   `json:"ttmSessionId"`
}

type commandResponse struct {
	Status

	CommandStatus string `json:"commandStatus"`
	Defaults      string `json:"defaults"`
	Groups        []int  `json:"groups"`
	SystemState   string `json:"systemState"`
}

func (api *Api) TurnAlarm(status bool) (bool, error) {
	systemStatus := "off"
	if status {
		systemStatus = "on"
	}
	emptyGroup := make([]string, 0)

	request := commandRequest{
		CurrentGroup: emptyGroup,
		NbGroups:     -1,
		SystemState:  systemStatus,
		TTMSessionId: api.context.TTMSessionId,
	}

	response, err := api.sendPost("/action/stateCommand", request)
	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(response.Body)
		println(string(bodyBytes))
		return false, errors.New("daitem server error")
	}

	var data commandResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return false, err
	}

	if err := checkValidStatus(data.Status); err != nil {
		return false, err
	}

	return true, nil
}
