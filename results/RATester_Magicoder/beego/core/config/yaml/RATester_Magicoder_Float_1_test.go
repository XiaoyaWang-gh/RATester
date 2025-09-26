package yaml

import (
	"errors"
	"fmt"
	"testing"
)

func TestFloat_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": float64(1.23),
			"key2": 456,
			"key3": int64(789),
			"key4": "not float64",
		},
	}

	tests := []struct {
		key      string
		expected float64
		err      error
	}{
		{"key1", 1.23, nil},
		{"key2", 456.0, nil},
		{"key3", 789.0, nil},
		{"key4", 0.0, errors.New("not float64 value")},
	}

	for _, test := range tests {
		got, err := cc.Float(test.key)
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Float(%q) returned error %v, want %v", test.key, err, test.err)
		}
		if got != test.expected {
			t.Errorf("Float(%q) returned %v, want %v", test.key, got, test.expected)
		}
	}
}
