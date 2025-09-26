package pagemeta

import (
	"fmt"
	"testing"
)

func TestIsZero_1(t *testing.T) {
	tests := []struct {
		name string
		b    BuildConfig
		want bool
	}{
		{
			name: "Test with zero value",
			b:    BuildConfig{},
			want: true,
		},
		{
			name: "Test with non-zero value",
			b: BuildConfig{
				List:             "always",
				Render:           "always",
				PublishResources: true,
				set:              true,
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

			if got := tt.b.IsZero(); got != tt.want {
				t.Errorf("BuildConfig.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
