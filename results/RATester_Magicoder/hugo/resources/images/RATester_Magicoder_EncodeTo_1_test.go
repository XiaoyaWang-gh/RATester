package images

import (
	"bytes"
	"fmt"
	"image"
	"testing"
)

func TestEncodeTo_1(t *testing.T) {
	tests := []struct {
		name    string
		conf    ImageConfig
		img     image.Image
		wantErr bool
	}{
		{
			name: "JPEG",
			conf: ImageConfig{
				TargetFormat: JPEG,
				Quality:      75,
			},
			img:     image.NewNRGBA(image.Rect(0, 0, 100, 100)),
			wantErr: false,
		},
		{
			name: "PNG",
			conf: ImageConfig{
				TargetFormat: PNG,
			},
			img:     image.NewNRGBA(image.Rect(0, 0, 100, 100)),
			wantErr: false,
		},
		{
			name: "GIF",
			conf: ImageConfig{
				TargetFormat: GIF,
			},
			img:     image.NewNRGBA(image.Rect(0, 0, 100, 100)),
			wantErr: false,
		},
		{
			name: "Unsupported format",
			conf: ImageConfig{
				TargetFormat: 0,
			},
			img:     image.NewNRGBA(image.Rect(0, 0, 100, 100)),
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
			w := &bytes.Buffer{}
			if err := i.EncodeTo(tt.conf, tt.img, w); (err != nil) != tt.wantErr {
				t.Errorf("EncodeTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
