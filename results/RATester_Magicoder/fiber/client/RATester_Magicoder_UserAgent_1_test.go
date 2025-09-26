package client

import (
	"fmt"
	"testing"
)

func TestUserAgent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &Request{
		userAgent: "TestUserAgent",
	}

	if req.UserAgent() != "TestUserAgent" {
		t.Errorf("Expected UserAgent to be 'TestUserAgent', but got '%s'", req.UserAgent())
	}
}
