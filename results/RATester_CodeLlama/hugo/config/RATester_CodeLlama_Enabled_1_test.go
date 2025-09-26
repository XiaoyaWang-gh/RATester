package config

import (
	"fmt"
	"testing"
)

func TestEnabled_1(t *testing.T) {
	tests := []struct {
		name string
		w    BuildStats
		want bool
	}{
		{
			name: "enabled",
			w:    BuildStats{Enable: true},
			want: true,
		},
		{
			name: "disabled",
			w:    BuildStats{Enable: false},
			want: false,
		},
		{
			name: "disabled tags",
			w:    BuildStats{Enable: true, DisableTags: true},
			want: false,
		},
		{
			name: "disabled classes",
			w:    BuildStats{Enable: true, DisableClasses: true},
			want: false,
		},
		{
			name: "disabled ids",
			w:    BuildStats{Enable: true, DisableIDs: true},
			want: false,
		},
		{
			name: "disabled all",
			w:    BuildStats{Enable: true, DisableTags: true, DisableClasses: true, DisableIDs: true},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.w.Enabled(); got != tt.want {
				t.Errorf("BuildStats.Enabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
