package json

import (
	"fmt"
	"testing"
)

func TestDefaultString_6(t *testing.T) {
	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key": "value",
		},
	}

	tests := []struct {
		name       string
		key        string
		defaultVal string
		want       string
	}{
		{
			name:       "Existing key",
			key:        "key",
			defaultVal: "default",
			want:       "value",
		},
		{
			name:       "Non-existing key",
			key:        "non-existing",
			defaultVal: "default",
			want:       "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := container.DefaultString(tt.key, tt.defaultVal); got != tt.want {
				t.Errorf("DefaultString() = %v, want %v", got, tt.want)
			}
		})
	}
}
