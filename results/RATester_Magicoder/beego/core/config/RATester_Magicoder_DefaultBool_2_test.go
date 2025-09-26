package config

import (
	"fmt"
	"testing"
)

func TestDefaultBool_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &fakeConfigContainer{
		data: map[string]string{
			"key1": "true",
			"key2": "false",
			"key3": "invalid",
		},
	}

	tests := []struct {
		key        string
		defaultVal bool
		expected   bool
	}{
		{"key1", false, true},
		{"key2", true, false},
		{"key3", true, true},
		{"key4", true, true},
	}

	for _, test := range tests {
		result := fc.DefaultBool(test.key, test.defaultVal)
		if result != test.expected {
			t.Errorf("DefaultBool(%s, %v) = %v, expected %v", test.key, test.defaultVal, result, test.expected)
		}
	}
}
