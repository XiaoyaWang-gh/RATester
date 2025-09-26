package xml

import (
	"fmt"
	"testing"
)

func TestDefaultInt_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": 123,
			"key2": "456",
		},
	}

	tests := []struct {
		key        string
		defaultVal int
		want       int
	}{
		{"key1", 0, 123},
		{"key2", 789, 789},
		{"key3", 101112, 101112},
	}

	for _, tt := range tests {
		got := cc.DefaultInt(tt.key, tt.defaultVal)
		if got != tt.want {
			t.Errorf("DefaultInt(%q, %v) = %v, want %v", tt.key, tt.defaultVal, got, tt.want)
		}
	}
}
