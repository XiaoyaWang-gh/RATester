package env

import (
	"fmt"
	"testing"
)

func TestMustGet_1(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		expected string
		err      error
	}{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: "value1",
			err:      nil,
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: "value2",
			err:      nil,
		},
		{
			name:     "Test case 3",
			key:      "key3",
			expected: "",
			err:      fmt.Errorf("no env variable with key3"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			env.Store(tt.key, tt.expected)
			val, err := MustGet(tt.key)
			if val != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, val)
			}
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("Expected error %v, got %v", tt.err, err)
			}
		})
	}
}
