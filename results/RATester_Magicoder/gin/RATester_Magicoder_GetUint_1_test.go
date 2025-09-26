package gin

import (
	"fmt"
	"testing"
)

func TestGetUint_1(t *testing.T) {
	type test struct {
		name     string
		key      string
		expected uint
	}

	tests := []test{
		{"TestGetUint_0", "key0", 1},
		{"TestGetUint_1", "key1", 2},
		{"TestGetUint_2", "key2", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Context{
				Keys: map[string]any{
					tt.key: tt.expected,
				},
			}

			got := c.GetUint(tt.key)
			if got != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, got)
			}
		})
	}
}
