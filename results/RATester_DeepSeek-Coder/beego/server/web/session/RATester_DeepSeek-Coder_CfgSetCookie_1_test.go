package session

import (
	"fmt"
	"testing"
)

func TestCfgSetCookie_1(t *testing.T) {
	t.Run("TestCfgSetCookie", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		config := &ManagerConfig{}
		CfgSetCookie(true)(config)
		if !config.EnableSetCookie {
			t.Errorf("Expected EnableSetCookie to be true, got false")
		}
	})
}
