package data

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAddDefaultHeaders_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &http.Request{}
	addDefaultHeaders(req, "application/json")
	if req.Header.Get("Accept") != "application/json" {
		t.Errorf("addDefaultHeaders failed")
	}
	if req.Header.Get("User-Agent") != "Hugo Static Site Generator" {
		t.Errorf("addDefaultHeaders failed")
	}
}
