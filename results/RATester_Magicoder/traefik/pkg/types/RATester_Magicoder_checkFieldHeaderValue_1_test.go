package types

import (
	"fmt"
	"testing"
)

func TestCheckFieldHeaderValue_1(t *testing.T) {
	tests := []struct {
		name         string
		value        string
		defaultValue string
		want         string
	}{
		{
			name:         "Keep",
			value:        "keep",
			defaultValue: "default",
			want:         "keep",
		},
		{
			name:         "Drop",
			value:        "drop",
			defaultValue: "default",
			want:         "drop",
		},
		{
			name:         "Redact",
			value:        "redact",
			defaultValue: "default",
			want:         "redact",
		},
		{
			name:         "Default",
			value:        "invalid",
			defaultValue: "default",
			want:         "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := checkFieldHeaderValue(tt.value, tt.defaultValue); got != tt.want {
				t.Errorf("checkFieldHeaderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
