package api

import (
	"github.com/stefanopulze/daitem/context"
	"net/http"
)

type Api struct {
	http    *http.Client
	context *context.Context
}

func NewApi(ctx *context.Context) *Api {
	return &Api{
		http:    &http.Client{},
		context: ctx,
	}
}
