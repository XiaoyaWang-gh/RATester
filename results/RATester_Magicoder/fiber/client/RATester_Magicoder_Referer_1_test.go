package client

import (
	"fmt"
	"testing"
)

func TestReferer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &Request{
		referer: "https://example.com",
	}

	expected := "https://example.com"
	actual := req.Referer()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
