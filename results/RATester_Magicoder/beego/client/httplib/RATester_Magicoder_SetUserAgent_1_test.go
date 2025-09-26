package httplib

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

	req := &BeegoHTTPRequest{
		setting: BeegoHTTPSettings{
			UserAgent: "TestUserAgent",
		},
	}

	newReq := req.SetUserAgent("NewUserAgent")

	if newReq.setting.UserAgent != "NewUserAgent" {
		t.Errorf("Expected UserAgent to be 'NewUserAgent', but got '%s'", newReq.setting.UserAgent)
	}
}
