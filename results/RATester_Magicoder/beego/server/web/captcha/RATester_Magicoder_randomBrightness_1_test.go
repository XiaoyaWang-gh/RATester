package captcha

import (
	"fmt"
	"image/color"
	"reflect"
	"testing"
)

func TestrandomBrightness_1(t *testing.T) {
	tests := []struct {
		name string
		c    color.RGBA
		max  uint8
		want color.RGBA
	}{
		{
			name: "Test case 1",
			c:    color.RGBA{R: 100, G: 100, B: 100, A: 255},
			max:  200,
			want: color.RGBA{R: 100, G: 100, B: 100, A: 255},
		},
		{
			name: "Test case 2",
			c:    color.RGBA{R: 150, G: 150, B: 150, A: 255},
			max:  200,
			want: color.RGBA{R: 150, G: 150, B: 150, A: 255},
		},
		{
			name: "Test case 3",
			c:    color.RGBA{R: 250, G: 250, B: 250, A: 255},
			max:  200,
			want: color.RGBA{R: 200, G: 200, B: 200, A: 255},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := randomBrightness(tt.c, tt.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randomBrightness() = %v, want %v", got, tt.want)
			}
		})
	}
}
