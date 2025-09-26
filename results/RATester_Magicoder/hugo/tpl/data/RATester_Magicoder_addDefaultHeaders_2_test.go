package data

import (
	"fmt"
	"net/http"
	"testing"
)

func TestaddDefaultHeaders_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	addDefaultHeaders(req, "application/json", "text/html")

	if req.Header.Get("Accept") != "application/json, text/html" {
		t.Error("Expected Accept header to be set")
	}

	if req.Header.Get("User-Agent") != "Hugo Static Site Generator" {
		t.Error("Expected User-Agent header to be set")
	}
}
