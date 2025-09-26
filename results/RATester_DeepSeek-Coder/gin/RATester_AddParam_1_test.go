package gin

import (
	"fmt"
	"testing"
)

func TestAddParam_1(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    string
		expected Param
	}{
		{
			name:     "Test case 1",
			key:      "key1",
			value:    "value1",
			expected: Param{Key: "key1", Value: "value1"},
		},
		{
			name:     "Test case 2",
			key:      "key2",
			value:    "value2",
			expected: Param{Key: "key2", Value: "value2"},
		},
		{
			name:     "Test case 3",
			key:      "key3",
			value:    "value3",
			expected: Param{Key: "key3", Value: "value3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Context{
				Params: make(Params, 0),
			}

			c.AddParam(tt.key, tt.value)

			if len(c.Params) != 1 {
				t.Errorf("Expected 1 param, got %d", len(c.Params))
			}

			param := c.Params[0]
			if param.Key != tt.expected.Key || param.Value != tt.expected.Value {
				t.Errorf("Expected param %v, got %v", tt.expected, param)
			}
		})
	}
}
