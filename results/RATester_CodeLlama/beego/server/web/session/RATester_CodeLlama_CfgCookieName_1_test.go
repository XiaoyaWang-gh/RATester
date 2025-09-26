package session

import (
	"fmt"
	"testing"
)

func TestCfgCookieName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var cookieName string
	var config ManagerConfig
	config.Opts(CfgCookieName(cookieName))
	if config.CookieName != cookieName {
		t.Errorf("config.CookieName = %v, want %v", config.CookieName, cookieName)
	}
}
