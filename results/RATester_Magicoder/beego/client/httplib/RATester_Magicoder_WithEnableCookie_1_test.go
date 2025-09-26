package httplib

import (
	"fmt"
	"testing"
)

func TestWithEnableCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		Name:     "test",
		Endpoint: "http://localhost",
		Setting: BeegoHTTPSettings{
			EnableCookie: false,
		},
	}

	WithEnableCookie(true)(client)

	if !client.Setting.EnableCookie {
		t.Error("EnableCookie should be true")
	}
}
