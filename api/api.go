package api

import (
	"fmt"
	"github.com/stefanopulze/daitem/http"
	"github.com/stefanopulze/daitem/session"
	goHttp "net/http"
)

type Client struct {
	http    *http.Client
	session *session.Session
}

func NewClient(http *http.Client, session *session.Session) *Client {
	return &Client{
		http:    http,
		session: session,
	}
}

func (c *Client) withBearerAuthorization() func(*goHttp.Request) {
	return func(request *goHttp.Request) {
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.session.GetAccessToken()))
	}
}
