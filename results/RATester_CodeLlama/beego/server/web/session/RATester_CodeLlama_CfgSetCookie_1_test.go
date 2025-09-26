package session

import (
	"fmt"
	"testing"
)

func TestCfgSetCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	enable := true
	opt := CfgSetCookie(enable)
	config := &ManagerConfig{}
	opt(config)
	if config.EnableSetCookie != enable {
		t.Errorf("config.EnableSetCookie = %v, want %v", config.EnableSetCookie, enable)
	}
}
