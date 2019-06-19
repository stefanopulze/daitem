package api

import (
	"encoding/json"
	"github.com/stefanopulze/daitem/data"
	"github.com/stefanopulze/daitem/errors"
	"time"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Status

	SessionId string `json:"sessionId"`
	Country   string `json:"country"`
}

func (api *Api) Login() (*data.DeviceSession, error) {
	if api.context.IsValid() {
		return &data.DeviceSession{
			SessionId:   api.context.SessionId,
			SessionTime: api.context.SessionTime,
			UseCache:    true,
		}, nil
	}

	request := loginRequest{
		Username: api.context.Username,
		Password: api.context.Password,
	}
	response, err := api.sendPost("/authenticate/login", request)

	if err != nil {
		return nil, errors.HttpError(err)
	}

	defer response.Body.Close()

	var respData loginResponse
	if err := json.NewDecoder(response.Body).Decode(&respData); err != nil {
		return nil, errors.JsonError(err)
	}

	if err := checkValidStatus(respData.Status); err != nil {
		return nil, err
	}

	api.context.SessionId = respData.SessionId
	api.context.SessionTime = time.Now()

	return &data.DeviceSession{
		SessionId:   api.context.SessionId,
		SessionTime: api.context.SessionTime,
		UseCache:    false,
	}, nil
}
