package config

import (
	"errors"
	"fmt"
	"testing"
)

func TestBool_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &IniConfigContainer{
		data: map[string]map[string]string{
			"section1": {
				"key1": "true",
				"key2": "false",
			},
		},
	}

	tests := []struct {
		key      string
		expected bool
		err      error
	}{
		{"section1:key1", true, nil},
		{"section1:key2", false, nil},
		{"section1:key3", false, errors.New("key not found")},
	}

	for _, test := range tests {
		val, err := container.Bool(test.key)
		if err != test.err {
			t.Errorf("Expected error %v, but got %v", test.err, err)
		}
		if val != test.expected {
			t.Errorf("Expected value %v, but got %v", test.expected, val)
		}
	}
}
