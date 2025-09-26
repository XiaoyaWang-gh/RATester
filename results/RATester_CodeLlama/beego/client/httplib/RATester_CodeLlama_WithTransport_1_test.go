package httplib

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWithTransport_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	transport := &http.Transport{}
	option := WithTransport(transport)
	client := &Client{}
	option(client)
	if client.Setting.Transport != transport {
		t.Errorf("WithTransport failed")
	}
}
