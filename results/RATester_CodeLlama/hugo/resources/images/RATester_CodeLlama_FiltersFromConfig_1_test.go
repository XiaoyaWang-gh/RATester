package images

import (
	"fmt"
	"image"
	"image/color"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestFiltersFromConfig_1(t *testing.T) {
	p := &ImageProcessor{}

	// TODO: Add more tests
	tests := []struct {
		name string
		src  image.Image
		conf ImageConfig
		want []gift.Filter
	}{
		{
			name: "rotate",
			src:  image.NewRGBA(image.Rect(0, 0, 100, 100)),
			conf: ImageConfig{
				Rotate: 90,
			},
			want: []gift.Filter{
				gift.Rotate(90, color.Transparent, gift.NearestNeighborInterpolation),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got, _ := p.FiltersFromConfig(tt.src, tt.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FiltersFromConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
