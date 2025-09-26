package yaml

import (
	"fmt"
	"testing"
)

func TestDefaultInt64_1(t *testing.T) {
	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": int64(10),
			"key2": "notInt64",
		},
	}

	tests := []struct {
		name       string
		key        string
		defaultVal int64
		want       int64
	}{
		{"existing key", "key1", 20, 10},
		{"non-existing key", "key3", 30, 30},
		{"non-int64 value", "key2", 40, 40},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := cc.DefaultInt64(tt.key, tt.defaultVal); got != tt.want {
				t.Errorf("DefaultInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
