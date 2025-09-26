package session

import (
	"fmt"
	"testing"
)

func TestCfgProviderConfig_1(t *testing.T) {
	testCases := []struct {
		name           string
		providerConfig string
		expected       string
	}{
		{
			name:           "Test case 1",
			providerConfig: "test_config_1",
			expected:       "test_config_1",
		},
		{
			name:           "Test case 2",
			providerConfig: "test_config_2",
			expected:       "test_config_2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			opt := CfgProviderConfig(tc.providerConfig)
			config := &ManagerConfig{}
			opt(config)
			if config.ProviderConfig != tc.expected {
				t.Errorf("Expected ProviderConfig to be %s, but got %s", tc.expected, config.ProviderConfig)
			}
		})
	}
}
