package api

import (
	"encoding/json"
	"errors"
	"fmt"
)

type isConnectedPayload struct {
	TransmitterId string `json:"transmitterId"`
}

type isConnectedResponse struct {
	IsConnected bool `json:"isConnected"`
}

type connectPayload struct {
	MasterCode string `json:"masterCode"`
}

func (c *Client) IsConnected(transmitterId string) (bool, error) {
	payload := isConnectedPayload{
		TransmitterId: transmitterId,
	}
	response, err := c.http.Post("/topaze/installation/isConnected", payload, c.withBearerAuthorization())
	if err != nil {
		return false, err
	}

	if response.StatusCode != 200 {
		return false, errors.New("Invalid response code: " + response.Status)
	}
	defer func() { _ = response.Body.Close() }()

	data := new(isConnectedResponse)
	if err = json.NewDecoder(response.Body).Decode(data); err != nil {
		return false, err
	}

	return data.IsConnected, nil
}

func (c *Client) Connect(configurationId int, masterCode string) (*SystemConnect, error) {
	url := fmt.Sprintf("/topaze/v5/systems/%d/connect", configurationId)
	payload := connectPayload{
		MasterCode: masterCode,
	}
	response, err := c.http.Post(url, payload, c.withBearerAuthorization())
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Invalid response code: " + response.Status)
	}
	defer func() { _ = response.Body.Close() }()

	data := new(SystemConnect)
	if err = json.NewDecoder(response.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}
