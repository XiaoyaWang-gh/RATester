package client

import (
	"fmt"
	"testing"
)

func TestURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &Request{
		url: "http://example.com",
	}

	expected := "http://example.com"
	actual := req.URL()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
