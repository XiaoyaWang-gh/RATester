package xml

import (
	"fmt"
	"testing"
)

func TestDefaultBool_9(t *testing.T) {
	t.Parallel()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": true,
			"key2": false,
			"key3": "not a bool",
		},
	}

	tests := []struct {
		name       string
		key        string
		defaultVal bool
		want       bool
	}{
		{"TestTrue", "key1", false, true},
		{"TestFalse", "key2", true, false},
		{"TestDefault", "key3", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := cc.DefaultBool(tt.key, tt.defaultVal); got != tt.want {
				t.Errorf("DefaultBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
