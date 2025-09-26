package create

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gohugoio/hugo/resources"
)

func TestFromRemote_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		rs:         &resources.Spec{},
		httpClient: &http.Client{},
	}

	uri := "http://example.com/resource"
	options := map[string]any{"method": "GET"}

	_, err := client.FromRemote(uri, options)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	_, err = client.FromRemote("invalid_url", options)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	_, err = client.FromRemote(uri, map[string]any{"invalid_option": "invalid_value"})
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
