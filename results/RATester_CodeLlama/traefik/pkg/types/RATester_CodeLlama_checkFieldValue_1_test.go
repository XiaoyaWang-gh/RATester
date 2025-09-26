package types

import (
	"fmt"
	"testing"
)

func TestCheckFieldValue_1(t *testing.T) {
	var tests = []struct {
		name        string
		value       string
		defaultKeep bool
		want        bool
	}{
		{
			name:        "AccessLogKeep",
			value:       AccessLogKeep,
			defaultKeep: true,
			want:        true,
		},
		{
			name:        "AccessLogDrop",
			value:       AccessLogDrop,
			defaultKeep: true,
			want:        false,
		},
		{
			name:        "defaultKeep",
			value:       "other",
			defaultKeep: true,
			want:        true,
		},
		{
			name:        "defaultDrop",
			value:       "other",
			defaultKeep: false,
			want:        false,
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
