package images

import (
	"fmt"
	"testing"
)

func TestSupportsTransparency_1(t *testing.T) {
	tests := []struct {
		name string
		f    Format
		want bool
	}{
		{
			name: "JPEG",
			f:    JPEG,
			want: false,
		},
		{
			name: "PNG",
			f:    2,
			want: true,
		},
		{
			name: "GIF",
			f:    3,
			want: true,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.f.SupportsTransparency(); got != tt.want {
				t.Errorf("Format.SupportsTransparency() = %v, want %v", got, tt.want)
			}
		})
	}
}
