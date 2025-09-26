package json

import (
	"fmt"
	"testing"
)

func TestBool_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key1": true,
			"key2": "false",
			"key3": 1,
		},
	}

	tests := []struct {
		key      string
		expected bool
		err      error
	}{
		{"key1", true, nil},
		{"key2", false, fmt.Errorf("not exist key: %q", "key2")},
		{"key3", false, fmt.Errorf("not exist key: %q", "key3")},
	}

	for _, test := range tests {
		val, err := container.Bool(test.key)
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error %v, got %v", test.err, err)
		}
		if val != test.expected {
			t.Errorf("Expected value %v, got %v", test.expected, val)
		}
	}
}
