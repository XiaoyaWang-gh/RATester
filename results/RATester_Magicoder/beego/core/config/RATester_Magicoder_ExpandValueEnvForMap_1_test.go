package config

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestExpandValueEnvForMap_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "test case 1",
			input: map[string]interface{}{
				"key1": "${ENV_VAR}",
				"key2": map[string]interface{}{
					"key3": "${ENV_VAR2}",
				},
				"key4": map[string]string{
					"key5": "${ENV_VAR3}",
				},
				"key6": map[interface{}]interface{}{
					"key7": "${ENV_VAR4}",
				},
			},
			expected: map[string]interface{}{
				"key1": "value1",
				"key2": map[string]interface{}{
					"key3": "value2",
				},
				"key4": map[string]string{
					"key5": "value3",
				},
				"key6": map[string]interface{}{
					"key7": "value4",
				},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			os.Setenv("ENV_VAR", "value1")
			os.Setenv("ENV_VAR2", "value2")
			os.Setenv("ENV_VAR3", "value3")
			os.Setenv("ENV_VAR4", "value4")

			result := ExpandValueEnvForMap(tc.input)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
