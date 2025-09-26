package captcha

import (
	"fmt"
	"image/color"
	"reflect"
	"testing"
)

func TestRandomPalette_1(t *testing.T) {
	tests := []struct {
		name string
		want color.Palette
	}{
		{
			name: "Test case 1",
			want: color.Palette{
				color.RGBA{0xFF, 0xFF, 0xFF, 0x00},
				color.RGBA{
					uint8(randIntn(129)),
					uint8(randIntn(129)),
					uint8(randIntn(129)),
					0xFF,
				},
				randomBrightness(color.RGBA{
					uint8(randIntn(129)),
					uint8(randIntn(129)),
					uint8(randIntn(129)),
					0xFF,
				}, 255),
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

			if got := randomPalette(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randomPalette() = %v, want %v", got, tt.want)
			}
		})
	}
}
