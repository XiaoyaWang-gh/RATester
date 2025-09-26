package main

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/plugins"
)

func TestCheckUniquePluginNames_1(t *testing.T) {
	type testCase struct {
		name     string
		exp      *static.Experimental
		expected error
	}

	testCases := []testCase{
		{
			name: "nil experimental",
			exp:  nil,
		},
		{
			name: "unique plugin names",
			exp: &static.Experimental{
				Plugins: map[string]plugins.Descriptor{
					"plugin1": {},
					"plugin2": {},
				},
				LocalPlugins: map[string]plugins.LocalDescriptor{
					"plugin3": {},
					"plugin4": {},
				},
			},
		},
		{
			name: "non-unique plugin names",
			exp: &static.Experimental{
				Plugins: map[string]plugins.Descriptor{
					"plugin1": {},
					"plugin2": {},
				},
				LocalPlugins: map[string]plugins.LocalDescriptor{
					"plugin1": {},
					"plugin3": {},
				},
			},
			expected: fmt.Errorf("the plugin's name %q must be unique", "plugin1"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := checkUniquePluginNames(test.exp)
			if err != nil && test.expected == nil {
				t.Errorf("Expected no error, got %v", err)
			} else if err == nil && test.expected != nil {
				t.Errorf("Expected error %v, got no error", test.expected)
			} else if err != nil && test.expected != nil && err.Error() != test.expected.Error() {
				t.Errorf("Expected error %v, got %v", test.expected, err)
			}
		})
	}
}
