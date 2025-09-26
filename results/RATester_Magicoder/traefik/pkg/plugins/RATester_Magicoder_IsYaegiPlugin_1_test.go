package plugins

import (
	"fmt"
	"testing"
)

func TestIsYaegiPlugin_1(t *testing.T) {
	tests := []struct {
		name string
		m    *Manifest
		want bool
	}{
		{
			name: "Test case 1",
			m: &Manifest{
				Runtime: "yaegi",
			},
			want: true,
		},
		{
			name: "Test case 2",
			m: &Manifest{
				Runtime: "",
			},
			want: true,
		},
		{
			name: "Test case 3",
			m: &Manifest{
				Runtime: "not_yaegi",
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

			if got := tt.m.IsYaegiPlugin(); got != tt.want {
				t.Errorf("IsYaegiPlugin() = %v, want %v", got, tt.want)
			}
		})
	}
}
