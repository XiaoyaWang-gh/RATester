package middleware

import (
	"errors"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestFindPluginConfig_1(t *testing.T) {
	tests := []struct {
		name        string
		rawConfig   map[string]dynamic.PluginConf
		expectedErr error
	}{
		{
			name: "single plugin",
			rawConfig: map[string]dynamic.PluginConf{
				"pluginType": {
					"field name": "value",
				},
			},
			expectedErr: nil,
		},
		{
			name: "multiple plugins",
			rawConfig: map[string]dynamic.PluginConf{
				"pluginType1": {
					"field name1": "value1",
				},
				"pluginType2": {
					"field name2": "value2",
				},
			},
			expectedErr: errors.New("invalid configuration: no configuration or too many plugin definition"),
		},
		{
			name:        "no plugins",
			rawConfig:   map[string]dynamic.PluginConf{},
			expectedErr: errors.New("missing plugin type"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, _, err := findPluginConfig(test.rawConfig)
			if err != nil && test.expectedErr == nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
			if err == nil && test.expectedErr != nil {
				t.Errorf("Expected error: %v, but got no error", test.expectedErr)
			}
			if err != nil && test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
				t.Errorf("Expected error: %v, but got: %v", test.expectedErr, err)
			}
		})
	}
}
