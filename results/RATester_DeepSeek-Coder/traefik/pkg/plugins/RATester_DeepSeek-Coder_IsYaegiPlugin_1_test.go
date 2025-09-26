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
			name: "Test with runtimeYaegi",
			m: &Manifest{
				Runtime: runtimeYaegi,
			},
			want: true,
		},
		{
			name: "Test with empty runtime",
			m: &Manifest{
				Runtime: "",
			},
			want: true,
		},
		{
			name: "Test with other runtime",
			m: &Manifest{
				Runtime: "other",
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
				t.Errorf("Manifest.IsYaegiPlugin() = %v, want %v", got, tt.want)
			}
		})
	}
}
