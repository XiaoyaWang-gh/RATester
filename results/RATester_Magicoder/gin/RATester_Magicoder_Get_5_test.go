package gin

import (
	"fmt"
	"testing"
)

func TestGet_5(t *testing.T) {
	params := Params{
		{"key1", "value1"},
		{"key2", "value2"},
	}

	tests := []struct {
		name     string
		params   Params
		key      string
		expected string
		found    bool
	}{
		{
			name:     "Found",
			params:   params,
			key:      "key1",
			expected: "value1",
			found:    true,
		},
		{
			name:     "Not Found",
			params:   params,
			key:      "key3",
			expected: "",
			found:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			value, found := test.params.Get(test.key)
			if value != test.expected || found != test.found {
				t.Errorf("Expected (%s, %t), got (%s, %t)", test.expected, test.found, value, found)
			}
		})
	}
}
