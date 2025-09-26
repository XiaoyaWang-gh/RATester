package proxy

import (
	"fmt"
	"testing"

	plugin "github.com/fatedier/frp/pkg/plugin/server"
)

func TestGetUserInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pxy := &BaseProxy{
		userInfo: plugin.UserInfo{
			User:  "testUser",
			Metas: map[string]string{"key": "value"},
			RunID: "testRunID",
		},
	}

	userInfo := pxy.GetUserInfo()

	if userInfo.User != "testUser" {
		t.Errorf("Expected user 'testUser', got '%s'", userInfo.User)
	}

	if userInfo.Metas["key"] != "value" {
		t.Errorf("Expected meta 'value', got '%s'", userInfo.Metas["key"])
	}

	if userInfo.RunID != "testRunID" {
		t.Errorf("Expected run ID 'testRunID', got '%s'", userInfo.RunID)
	}
}
