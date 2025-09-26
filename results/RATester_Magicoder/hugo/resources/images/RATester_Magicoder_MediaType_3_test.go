package images

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/media"
)

func TestMediaType_3(t *testing.T) {
	tests := []struct {
		name string
		f    Format
		want media.Type
	}{
		{
			name: "JPEG",
			f:    JPEG,
			want: media.Builtin.JPEGType,
		},
		{
			name: "PNG",
			f:    PNG,
			want: media.Builtin.PNGType,
		},
		{
			name: "GIF",
			f:    GIF,
			want: media.Builtin.GIFType,
		},
		{
			name: "TIFF",
			f:    TIFF,
			want: media.Builtin.TIFFType,
		},
		{
			name: "BMP",
			f:    BMP,
			want: media.Builtin.BMPType,
		},
		{
			name: "WEBP",
			f:    WEBP,
			want: media.Builtin.WEBPType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.f.MediaType(); got != tt.want {
				t.Errorf("MediaType() = %v, want %v", got, tt.want)
			}
		})
	}
}
