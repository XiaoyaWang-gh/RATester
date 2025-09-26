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
			name: "all enabled",
			w:    BuildStats{Enable: true, DisableTags: false, DisableClasses: false, DisableIDs: false},
			want: true,
		},
		{
			name: "all disabled",
			w:    BuildStats{Enable: false, DisableTags: true, DisableClasses: true, DisableIDs: true},
			want: false,
		},
		{
			name: "some enabled",
			w:    BuildStats{Enable: true, DisableTags: true, DisableClasses: false, DisableIDs: true},
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
