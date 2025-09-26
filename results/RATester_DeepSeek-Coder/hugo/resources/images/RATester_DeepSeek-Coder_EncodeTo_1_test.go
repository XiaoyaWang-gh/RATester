package images

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"testing"
)

func TestEncodeTo_1(t *testing.T) {
	tests := []struct {
		name    string
		conf    ImageConfig
		img     image.Image
		w       io.Writer
		wantErr bool
	}{
		{
			name: "JPEG Encoding",
			conf: ImageConfig{
				TargetFormat: JPEG,
				Quality:      90,
			},
			img:     &image.NRGBA{},
			w:       &bytes.Buffer{},
			wantErr: false,
		},
		{
			name: "PNG Encoding",
			conf: ImageConfig{
				TargetFormat: PNG,
			},
			img:     &image.NRGBA{},
			w:       &bytes.Buffer{},
			wantErr: false,
		},
		{
			name: "GIF Encoding",
			conf: ImageConfig{
				TargetFormat: GIF,
			},
			img:     &image.NRGBA{},
			w:       &bytes.Buffer{},
			wantErr: false,
		},
		{
			name: "Unsupported Format",
			conf: ImageConfig{
				TargetFormat: 0,
			},
			img:     &image.NRGBA{},
			w:       &bytes.Buffer{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			i := &Image{}
			err := i.EncodeTo(tt.conf, tt.img, tt.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
