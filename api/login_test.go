package api

import (
	"github.com/stefanopulze/daitem/context"
	"github.com/stefanopulze/daitem/errors"
	"log"
	"testing"
)

func TestApi_Login(t *testing.T) {
	context := context.Context{}
	api := NewApi(&context)

	_, e := api.Login()

	if e == nil {
		t.Fatal("Error must be throw if context is expired")
	}

	log.Println(e)

	if de, ok := e.(*errors.ApiError); ok {
		if de.Code != 1 {
			t.Fatal("Context expired, code must be: 1")
		}
	}
}

func TestApi_LoginWithBadCredential(t *testing.T) {
	context := context.Context{
		Username: "john",
		Password: "doe",
	}
	api := NewApi(&context)

	_, e := api.Login()

	if e == nil {
		t.Fatal("Error must be throw if context is expired")
	}

	log.Println(e)

	if de, ok := e.(*errors.ApiError); ok {
		if de.Code != errors.HTTP {
			t.Fatal("Context expired, code must be: 1")
		}
	}
}
