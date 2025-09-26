package config

import (
	"fmt"
	"testing"
)

func TestDefaultInt_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &fakeConfigContainer{
		data: map[string]string{
			"key1": "10",
			"key2": "20",
		},
	}

	tests := []struct {
		key        string
		defaultVal int
		expected   int
	}{
		{"key1", 0, 10},
		{"key2", 0, 20},
		{"key3", 30, 30},
	}

	for _, test := range tests {
		result := fc.DefaultInt(test.key, test.defaultVal)
		if result != test.expected {
			t.Errorf("DefaultInt(%s, %d) = %d; want %d", test.key, test.defaultVal, result, test.expected)
		}
	}
}
