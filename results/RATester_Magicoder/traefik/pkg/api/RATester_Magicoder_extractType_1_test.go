package api

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestextractType_1(t *testing.T) {
	type testStruct struct {
		field1 map[string]dynamic.PluginConf
		field2 *struct{}
	}

	testCases := []struct {
		name     string
		input    testStruct
		expected string
	}{
		{
			name: "Test case 1",
			input: testStruct{
				field1: map[string]dynamic.PluginConf{
					"test": {},
				},
				field2: nil,
			},
			expected: "field1",
		},
		{
			name: "Test case 2",
			input: testStruct{
				field1: nil,
				field2: &struct{}{},
			},
			expected: "field2",
		},
		{
			name: "Test case 3",
			input: testStruct{
				field1: nil,
				field2: nil,
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

			result := extractType(&tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
