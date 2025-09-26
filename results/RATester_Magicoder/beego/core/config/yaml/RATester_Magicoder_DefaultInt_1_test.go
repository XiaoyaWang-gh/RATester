package yaml

import (
	"fmt"
	"testing"
)

func TestDefaultInt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": 123,
			"key2": "abc",
		},
	}

	tests := []struct {
		key        string
		defaultVal int
		want       int
	}{
		{"key1", 456, 123},
		{"key2", 789, 456},
		{"key3", 101112, 101112},
	}

	for _, tt := range tests {
		got := cc.DefaultInt(tt.key, tt.defaultVal)
		if got != tt.want {
			t.Errorf("DefaultInt(%q, %v) = %v, want %v", tt.key, tt.defaultVal, got, tt.want)
		}
	}
}
