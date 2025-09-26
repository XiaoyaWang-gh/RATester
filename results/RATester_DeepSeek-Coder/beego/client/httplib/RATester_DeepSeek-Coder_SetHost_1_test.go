package httplib

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetHost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &BeegoHTTPRequest{
		req: &http.Request{
			Host: "old.host",
		},
	}

	newHost := "new.host"
	req.SetHost(newHost)

	if req.req.Host != newHost {
		t.Errorf("Expected host to be %s, got %s", newHost, req.req.Host)
	}
}
