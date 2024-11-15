package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Status struct {
	Status  string `json:"status"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func (api *Api) sendPost(url string, data interface{}) (*http.Response, error) {
	fullUrl := baseUrl + url

	if jsonString, err := json.Marshal(data); err != nil {
		return nil, err
	} else {
		return api.http.Post(fullUrl, contentType, bytes.NewBuffer(jsonString))
	}
}

func checkValidStatus(response interface{}) error {
	// TODO: leggere i messaggi di errore e dare una costante ad ognuno
	if status, ok := response.(Status); ok {
		if status.Status == "KO" {
			message := "invalid api response: "
			if status.Error != "" {
				message += status.Error
			} else if status.Message != "" {
				message += status.Message
			} else {
				message += "generic error"
			}

			log.Printf(message)
			return errors.New(message)
		}
	}

	return nil
}
