package types

import (
	"fmt"
	"testing"
)

func TestCheckFieldValue_1(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		defaultKeep bool
		want        bool
	}{
		{
			name:        "Keep",
			value:       AccessLogKeep,
			defaultKeep: true,
			want:        true,
		},
		{
			name:        "Drop",
			value:       AccessLogDrop,
			defaultKeep: false,
			want:        false,
		},
		{
			name:        "Default",
			value:       "default",
			defaultKeep: true,
			want:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := checkFieldValue(tt.value, tt.defaultKeep); got != tt.want {
				t.Errorf("checkFieldValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
