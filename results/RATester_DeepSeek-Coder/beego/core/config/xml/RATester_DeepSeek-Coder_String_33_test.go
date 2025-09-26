package xml

import (
	"fmt"
	"testing"
)

func TestString_33(t *testing.T) {
	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": "value1",
			"key2": 123,
		},
	}

	tests := []struct {
		name     string
		key      string
		expected string
		err      error
	}{
		{"existing string", "key1", "value1", nil},
		{"non-existing string", "key2", "", nil},
		{"non-existing key", "key3", "", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := cc.String(tt.key)
			if err != tt.err {
				t.Errorf("expected error %v, got %v", tt.err, err)
			}
			if got != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}
