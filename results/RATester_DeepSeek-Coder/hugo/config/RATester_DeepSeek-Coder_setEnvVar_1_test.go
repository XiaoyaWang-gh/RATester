package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetEnvVar_1(t *testing.T) {
	testCases := []struct {
		name     string
		vars     []string
		key      string
		value    string
		expected []string
	}{
		{
			name:     "Test case 1",
			vars:     []string{"key1=value1", "key2=value2"},
			key:      "key1",
			value:    "newValue1",
			expected: []string{"key1=newValue1", "key2=value2"},
		},
		{
			name:     "Test case 2",
			vars:     []string{"key1=value1", "key2=value2"},
			key:      "key3",
			value:    "newValue3",
			expected: []string{"key1=value1", "key2=value2", "key3=newValue3"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			setEnvVar(&tc.vars, tc.key, tc.value)
			if !reflect.DeepEqual(tc.vars, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, tc.vars)
			}
		})
	}
}
