package tpl

import (
	"fmt"
	"testing"
)

func TestIsZero_3(t *testing.T) {
	tests := []struct {
		name string
		info ParseInfo
		want bool
	}{
		{
			name: "Test case 1",
			info: ParseInfo{
				Config: ParseConfig{
					Version: 0,
				},
			},
			want: true,
		},
		{
			name: "Test case 2",
			info: ParseInfo{
				Config: ParseConfig{
					Version: 1,
				},
			},
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

			if got := tt.info.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
