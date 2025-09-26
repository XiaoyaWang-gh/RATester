package config

import (
	"errors"
	"fmt"
	"testing"
)

func TestBool_3(t *testing.T) {
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
		key      string
		expected bool
		err      error
	}{
		{"key1", true, nil},
		{"key2", false, nil},
		{"key3", false, errors.New("invalid syntax")},
	}

	for _, test := range tests {
		val, err := fc.Bool(test.key)
		if err != test.err {
			t.Errorf("expected error %v, got %v", test.err, err)
		}
		if val != test.expected {
			t.Errorf("expected value %v, got %v", test.expected, val)
		}
	}
}
