package client

import (
	"fmt"
	"testing"
)

func TestSetBaseURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}
	url := "http://example.com"
	client.SetBaseURL(url)
	if client.baseURL != url {
		t.Errorf("Expected baseURL to be %s, but got %s", url, client.baseURL)
	}
}
