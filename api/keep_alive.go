package api

import "encoding/json"

type keepAliveRequest struct {
	SessionId    string `json:"sessionId"`
	TTMSessionId string `json:"ttmSessionId"`
}

func (api *Api) KeepAlive() error {
	request := keepAliveRequest{
		SessionId:    api.context.SessionId,
		TTMSessionId: api.context.TTMSessionId,
	}

	response, err := api.sendPost("/authenticate/keepAlive", request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	var status Status

	if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
		return err
	}

	if err := checkValidStatus(status); err != nil {
		return err
	}

	return nil
}
