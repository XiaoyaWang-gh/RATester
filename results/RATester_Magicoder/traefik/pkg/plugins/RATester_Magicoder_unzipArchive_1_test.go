package plugins

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestUnzipArchive_1(t *testing.T) {
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

	// Test case 1: Successful unzip
	err := client.unzipArchive("test", "1.0.0")
	if err != nil {
		t.Errorf("unzipArchive() failed: %v", err)
	}

	// Test case 2: Unzip with error
	err = client.unzipArchive("test", "invalid")
	if err == nil {
		t.Error("unzipArchive() should have failed with invalid version")
	}
}
