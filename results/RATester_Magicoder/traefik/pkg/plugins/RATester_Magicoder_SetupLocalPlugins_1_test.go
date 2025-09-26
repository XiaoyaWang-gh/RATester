package plugins

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestSetupLocalPlugins_1(t *testing.T) {
	tests := []struct {
		name     string
		plugins  map[string]LocalDescriptor
		expected error
	}{
		{
			name:     "nil plugins",
			plugins:  nil,
			expected: nil,
		},
		{
			name:     "empty plugins",
			plugins:  map[string]LocalDescriptor{},
			expected: nil,
		},
		{
			name: "valid plugins",
			plugins: map[string]LocalDescriptor{
				"plugin1": {ModuleName: "valid_plugin"},
				"plugin2": {ModuleName: "valid_plugin2"},
			},
			expected: nil,
		},
		{
			name: "invalid plugins",
			plugins: map[string]LocalDescriptor{
				"plugin1": {ModuleName: ""},
				"plugin2": {ModuleName: "/invalid_plugin"},
				"plugin3": {ModuleName: "invalid_plugin/"},
				"plugin4": {ModuleName: "duplicate_plugin"},
				"plugin5": {ModuleName: "duplicate_plugin"},
			},
			expected: errors.New("1 error occurred:\n\t* plugin1: plugin name is missing\n\t* plugin2: plugin name should not start or end with a /\n\t* only one version of a plugin is allowed, there is a duplicate of duplicate_plugin\n\n"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := SetupLocalPlugins(test.plugins)
			if !reflect.DeepEqual(err, test.expected) {
				t.Errorf("Expected error: %v, got: %v", test.expected, err)
			}
		})
	}
}
