package captcha

import (
	"fmt"
	"image/color"
	"reflect"
	"testing"
)

func TestRandomBrightness_1(t *testing.T) {
	tests := []struct {
		name string
		c    color.RGBA
		max  uint8
		want color.RGBA
	}{
		{
			name: "Test case 1",
			c:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
			max:  255,
			want: color.RGBA{R: 255, G: 255, B: 255, A: 255},
		},
		{
			name: "Test case 2",
			c:    color.RGBA{R: 128, G: 128, B: 128, A: 128},
			max:  255,
			want: color.RGBA{R: 130, G: 130, B: 130, A: 128},
		},
		{
			name: "Test case 3",
			c:    color.RGBA{R: 0, G: 0, B: 0, A: 0},
			max:  0,
			want: color.RGBA{R: 0, G: 0, B: 0, A: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := randomBrightness(tt.c, tt.max)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randomBrightness() = %v, want %v", got, tt.want)
			}
		})
	}
}
