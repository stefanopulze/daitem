package api

import "encoding/json"

type disconnectRequest struct {
	SessionId    string  `json:"sessionId"`
	TtmSessionId *string `json:"ttmSessionId"`
}

func (api *Api) Disconnect() error {
	request := disconnectRequest{
		SessionId: api.context.SessionId,
	}

	if api.context.TTMSessionId != "" {
		request.TtmSessionId = &api.context.TTMSessionId
	}

	response, err := api.sendPost("/authenticate/disconnect", request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	var data Status
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return err
	}

	if err := checkValidStatus(data); err != nil {
		return err
	}

	// TODO resettare sessionId e TTMSession?

	return nil
}
