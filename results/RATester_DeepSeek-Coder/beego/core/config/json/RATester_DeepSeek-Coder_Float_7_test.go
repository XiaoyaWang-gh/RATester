package json

import (
	"errors"
	"fmt"
	"testing"
)

func TestFloat_7(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		key      string
		expected float64
		err      error
	}{
		{
			name:     "test float64",
			key:      "key1",
			expected: 1.1,
			err:      nil,
		},
		{
			name:     "test not exist key",
			key:      "not_exist_key",
			expected: 0.0,
			err:      errors.New("not exist key:not_exist_key"),
		},
		{
			name:     "test not float64 value",
			key:      "key2",
			expected: 0.0,
			err:      errors.New("not float64 value"),
		},
	}

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key1": 1.1,
			"key2": "not float64",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			val, err := container.Float(tt.key)
			if val != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, val)
			}
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("expected error %v, got %v", tt.err, err)
			}
		})
	}
}
