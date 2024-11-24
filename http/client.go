package http

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Client is a wrap of native net/http client with some useful functions
type Client struct {
	baseUrl string
	headers map[string]string
}

type WithOption = func(*http.Request)

func NewClient(url string) *Client {
	return &Client{
		baseUrl: url,
		headers: make(map[string]string),
	}
}

func (c *Client) AddHeader(key, value string) {
	c.headers[key] = value
}

// Get send GET request
func (c *Client) Get(url string, options ...WithOption) (*http.Response, error) {
	request, err := http.NewRequest("GET", c.baseUrl+url, nil)
	if err != nil {
		return nil, err
	}

	c.addDefaultHeaders(request)

	for _, option := range options {
		option(request)
	}

	return http.DefaultClient.Do(request)
}

// Post send POST request
func (c *Client) Post(url string, payload interface{}, options ...WithOption) (*http.Response, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", c.baseUrl+url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	c.addDefaultHeaders(request)

	for _, option := range options {
		option(request)
	}

	return http.DefaultClient.Do(request)
}

func (c *Client) addDefaultHeaders(request *http.Request) {
	for key, value := range c.headers {
		request.Header.Add(key, value)
	}
}
