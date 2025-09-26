package json

import (
	"errors"
	"fmt"
	"testing"
)

func TestFloat_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key1": float64(1.23),
			"key2": "not float64",
			"key3": nil,
		},
	}

	tests := []struct {
		key      string
		expected float64
		err      error
	}{
		{"key1", 1.23, nil},
		{"key2", 0.0, errors.New("not float64 value")},
		{"key3", 0.0, errors.New("not exist key:key3")},
	}

	for _, test := range tests {
		val, err := container.Float(test.key)
		if val != test.expected || err != nil {
			t.Errorf("Expected (%v, %v), got (%v, %v)", test.expected, test.err, val, err)
		}
	}
}
