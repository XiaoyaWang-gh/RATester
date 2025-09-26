package session

import (
	"fmt"
	"testing"
)

func TestCfgHTTPOnly_1(t *testing.T) {
	t.Run("TestCfgHTTPOnly", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		config := &ManagerConfig{}
		CfgHTTPOnly(true)(config)
		if config.DisableHTTPOnly != false {
			t.Errorf("Expected DisableHTTPOnly to be false, got %v", config.DisableHTTPOnly)
		}
		CfgHTTPOnly(false)(config)
		if config.DisableHTTPOnly != true {
			t.Errorf("Expected DisableHTTPOnly to be true, got %v", config.DisableHTTPOnly)
		}
	})
}
