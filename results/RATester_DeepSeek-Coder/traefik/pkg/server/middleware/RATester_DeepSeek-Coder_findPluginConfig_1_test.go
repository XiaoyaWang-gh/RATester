package middleware

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestFindPluginConfig_1(t *testing.T) {
	type testCase struct {
		name          string
		rawConfig     map[string]dynamic.PluginConf
		expectedType  string
		expectedError bool
	}

	testCases := []testCase{
		{
			name: "One plugin",
			rawConfig: map[string]dynamic.PluginConf{
				"pluginType1": {},
			},
			expectedType:  "pluginType1",
			expectedError: false,
		},
		{
			name: "No plugin",
			rawConfig: map[string]dynamic.PluginConf{
				"": {},
			},
			expectedType:  "",
			expectedError: true,
		},
		{
			name: "Two plugins",
			rawConfig: map[string]dynamic.PluginConf{
				"pluginType1": {},
				"pluginType2": {},
			},
			expectedType:  "",
			expectedError: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			pluginType, _, err := findPluginConfig(test.rawConfig)

			if (err != nil) != test.expectedError {
				t.Errorf("Error: got %v, want %v", err, test.expectedError)
			}

			if pluginType != test.expectedType {
				t.Errorf("Plugin type: got %v, want %v", pluginType, test.expectedType)
			}
		})
	}
}
