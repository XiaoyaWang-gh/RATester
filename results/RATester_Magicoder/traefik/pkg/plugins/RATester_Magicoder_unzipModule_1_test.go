package plugins

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestUnzipModule_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		HTTPClient: &http.Client{},
		baseURL:    &url.URL{},
		archives:   "/tmp/archives",
		stateFile:  "/tmp/state.json",
		goPath:     "/tmp/go",
		sources:    "/tmp/sources",
	}

	pName := "test"
	pVersion := "1.0.0"

	err := client.unzipModule(pName, pVersion)
	if err != nil {
		t.Errorf("unzipModule returned an error: %v", err)
	}
}
