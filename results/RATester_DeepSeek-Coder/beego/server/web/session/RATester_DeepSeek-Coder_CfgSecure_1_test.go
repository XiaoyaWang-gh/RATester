package session

import (
	"fmt"
	"testing"
)

func TestCfgSecure_1(t *testing.T) {
	t.Run("TestCfgSecure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cfg := &ManagerConfig{}
		CfgSecure(true)(cfg)
		if !cfg.Secure {
			t.Errorf("Expected Secure to be true, got %v", cfg.Secure)
		}
	})
}
