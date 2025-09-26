package images

import (
	"fmt"
	"testing"
)

func TestRequiresDefaultQuality_1(t *testing.T) {
	tests := []struct {
		name string
		f    Format
		want bool
	}{
		{
			name: "JPEG",
			f:    JPEG,
			want: true,
		},
		{
			name: "WEBP",
			f:    WEBP,
			want: true,
		},
		{
			name: "PNG",
			f:    2,
			want: false,
		},
		{
			name: "BMP",
			f:    3,
			want: false,
		},
		{
			name: "GIF",
			f:    4,
			want: false,
		},
		{
			name: "TIFF",
			f:    5,
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

			if got := tt.f.RequiresDefaultQuality(); got != tt.want {
				t.Errorf("Format.RequiresDefaultQuality() = %v, want %v", got, tt.want)
			}
		})
	}
}
