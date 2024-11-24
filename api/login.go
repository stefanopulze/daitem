package api

import (
	"encoding/json"
	"errors"
)

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
	DiagralId    string `json:"diagralId"`
	UserId       int    `json:"userId"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

func (c *Client) AuthWithCredentials(username string, password string) (*LoginResponse, error) {
	payload := loginRequest{
		Username: username,
		Password: password,
	}

	response, err := c.http.Post("/topaze/v2/authenticate/login", payload)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, ErrInvalidCredentials
	}

	defer func() {
		_ = response.Body.Close()
	}()

	data := new(LoginResponse)
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) AuthWithRefreshToken(refreshToken string) (*LoginResponse, error) {
	payload := refreshTokenRequest{
		RefreshToken: refreshToken,
	}
	response, err := c.http.Post("topaze/v1/authenticate/refresh", payload)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, ErrInvalidRefreshToken
	}

	defer func() {
		_ = response.Body.Close()
	}()

	data := new(LoginResponse)
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
