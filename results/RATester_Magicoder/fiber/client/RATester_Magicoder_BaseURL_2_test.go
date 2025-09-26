package client

import (
	"fmt"
	"testing"
)

func TestBaseURL_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		baseURL: "https://example.com",
	}

	expected := "https://example.com"
	actual := client.BaseURL()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
