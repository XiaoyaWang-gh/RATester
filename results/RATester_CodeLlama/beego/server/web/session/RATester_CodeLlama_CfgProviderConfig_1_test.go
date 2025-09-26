package session

import (
	"fmt"
	"testing"
)

func TestCfgProviderConfig_1(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		providerConfig := "test"
		config := &ManagerConfig{}
		CfgProviderConfig(providerConfig)(config)
		if config.ProviderConfig != providerConfig {
			t.Errorf("expected %s, got %s", providerConfig, config.ProviderConfig)
		}
	})
}
