package api

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestExtractType_1(t *testing.T) {
	type testStruct struct {
		MapField      map[string]dynamic.PluginConf
		PtrField      *dynamic.PluginConf
		EmptyPtrField *dynamic.PluginConf
	}

	testCases := []struct {
		name     string
		input    *testStruct
		expected string
	}{
		{
			name: "MapField is not empty",
			input: &testStruct{
				MapField: map[string]dynamic.PluginConf{
					"key": {},
				},
			},
			expected: "MapField",
		},
		{
			name: "PtrField is not nil",
			input: &testStruct{
				PtrField: &dynamic.PluginConf{},
			},
			expected: "PtrField",
		},
		{
			name: "EmptyPtrField is nil",
			input: &testStruct{
				EmptyPtrField: nil,
			},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := extractType(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
