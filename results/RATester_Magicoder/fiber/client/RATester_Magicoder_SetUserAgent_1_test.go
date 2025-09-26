package client

import (
	"fmt"
	"testing"
)

func TestSetUserAgent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &Request{}
	ua := "TestUserAgent"
	req.SetUserAgent(ua)
	if req.userAgent != ua {
		t.Errorf("Expected user agent to be %s, but got %s", ua, req.userAgent)
	}
}
