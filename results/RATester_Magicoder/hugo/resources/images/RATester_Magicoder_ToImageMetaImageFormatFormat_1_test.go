package images

import (
	"fmt"
	"testing"

	"github.com/bep/imagemeta"
)

func TestToImageMetaImageFormatFormat_1(t *testing.T) {
	tests := []struct {
		name string
		f    Format
		want imagemeta.ImageFormat
	}{
		{
			name: "JPEG",
			f:    JPEG,
			want: imagemeta.JPEG,
		},
		{
			name: "PNG",
			f:    PNG,
			want: imagemeta.PNG,
		},
		{
			name: "TIFF",
			f:    TIFF,
			want: imagemeta.TIFF,
		},
		{
			name: "WEBP",
			f:    WEBP,
			want: imagemeta.WebP,
		},
		{
			name: "Unknown",
			f:    -1,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.f.ToImageMetaImageFormatFormat(); got != tt.want {
				t.Errorf("ToImageMetaImageFormatFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
