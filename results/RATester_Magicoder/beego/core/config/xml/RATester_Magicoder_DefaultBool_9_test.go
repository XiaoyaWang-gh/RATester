package xml

import (
	"fmt"
	"testing"
)

func TestDefaultBool_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": true,
			"key2": false,
			"key3": "not a bool",
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
		{"key4", false, false},
	}

	for _, test := range tests {
		result := cc.DefaultBool(test.key, test.defaultVal)
		if result != test.expected {
			t.Errorf("DefaultBool(%s, %v) = %v; want %v", test.key, test.defaultVal, result, test.expected)
		}
	}
}
