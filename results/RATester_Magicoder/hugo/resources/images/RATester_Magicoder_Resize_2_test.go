package images

import (
	"fmt"
	"image"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestResize_2(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	resizer := imagingResizer{
		p:      &ImageProcessor{},
		filter: gift.LanczosResampling,
	}

	tests := []struct {
		name   string
		width  uint
		height uint
		want   image.Image
	}{
		{
			name:   "Resize to 50x50",
			width:  50,
			height: 50,
			want:   image.NewRGBA(image.Rect(0, 0, 50, 50)),
		},
		{
			name:   "Resize to 0x50",
			width:  0,
			height: 50,
			want:   image.NewRGBA(image.Rect(0, 0, 20, 50)),
		},
		{
			name:   "Resize to 50x0",
			width:  50,
			height: 0,
			want:   image.NewRGBA(image.Rect(0, 0, 50, 20)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := resizer.Resize(img, tt.width, tt.height)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resize() = %v, want %v", got, tt.want)
			}
		})
	}
}
