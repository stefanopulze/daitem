package api

import (
	"encoding/json"
	"errors"
	"fmt"
)

type sendCommandPayload struct {
	Active bool `json:"active"`
}

type getConfigurationPayload struct {
	SystemId int `json:"systemId"`
	Role     int `json:"role"`
}

// Systems List all user's systems
func (c *Client) Systems() ([]System, error) {
	response, err := c.http.Get("/topaze/v5/systems", c.withBearerAuthorization())
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Invalid response code: " + response.Status)
	}
	defer func() { _ = response.Body.Close() }()

	systems := make([]System, 0)
	if err = json.NewDecoder(response.Body).Decode(&systems); err != nil {
		return nil, err
	}

	return systems, nil
}

func (c *Client) GetSystemConfiguration(systemId int) (*Configuration, error) {
	payload := getConfigurationPayload{
		SystemId: systemId,
		Role:     1,
	}

	response, err := c.http.Post("/topaze/configuration/getConfiguration", payload, c.withBearerAuthorization())
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Invalid response code: " + response.Status)
	}
	defer func() { _ = response.Body.Close() }()

	configuration := new(Configuration)
	if err = json.NewDecoder(response.Body).Decode(configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

// SystemState get current system state
func (c *Client) SystemState(systemId int) (*SystemState, error) {
	url := fmt.Sprintf("/topaze/v5/systems/%d/state", systemId)
	response, err := c.http.Get(url, c.withBearerAuthorization())
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Invalid response code: " + response.Status)
	}
	defer func() { _ = response.Body.Close() }()

	data := new(SystemState)
	if err = json.NewDecoder(response.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}

// SystemSendCommand send active or deactivate command to system
func (c *Client) SystemSendCommand(systemId int, active bool) (*SystemState, error) {
	url := fmt.Sprintf("/topaze/v1/action/systems/%d/sendSystemCommand", systemId)
	payload := sendCommandPayload{Active: active}
	response, err := c.http.Post(url, payload, c.withBearerAuthorization())
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Invalid response code: " + response.Status)
	}
	defer func() { _ = response.Body.Close() }()

	data := new(SystemState)
	if err = json.NewDecoder(response.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}

// SystemRefresh get information about system
func (c *Client) SystemRefresh(systemId int) (*SystemInfo, error) {
	url := fmt.Sprintf("/topaze/v5/systems/%d/refresh", systemId)
	response, err := c.http.Post(url, nil, c.withBearerAuthorization())
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Invalid response code: " + response.Status)
	}
	defer func() { _ = response.Body.Close() }()

	data := new(SystemInfo)
	if err = json.NewDecoder(response.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}
