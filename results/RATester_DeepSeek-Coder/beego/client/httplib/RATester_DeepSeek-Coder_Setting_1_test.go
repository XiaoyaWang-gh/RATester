package httplib

import (
	"fmt"
	"testing"
)

func TestSetting_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{}
	setting := BeegoHTTPSettings{
		UserAgent: "TestUserAgent",
	}
	b.Setting(setting)
	if b.setting.UserAgent != "TestUserAgent" {
		t.Errorf("Expected UserAgent to be 'TestUserAgent', got '%s'", b.setting.UserAgent)
	}
}
