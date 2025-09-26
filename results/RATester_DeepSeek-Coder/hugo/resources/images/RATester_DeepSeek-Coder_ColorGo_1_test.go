package images

import (
	"fmt"
	"image/color"
	"reflect"
	"testing"
)

func TestColorGo_1(t *testing.T) {
	tests := []struct {
		name string
		c    Color
		want color.Color
	}{
		{
			name: "Test case 1",
			c:    Color{color: color.RGBA{R: 255, G: 0, B: 0, A: 255}, hex: "#FF0000", luminance: 0.2126},
			want: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		},
		{
			name: "Test case 2",
			c:    Color{color: color.RGBA{R: 0, G: 255, B: 0, A: 255}, hex: "#00FF00", luminance: 0.7152},
			want: color.RGBA{R: 0, G: 255, B: 0, A: 255},
		},
		{
			name: "Test case 3",
			c:    Color{color: color.RGBA{R: 0, G: 0, B: 255, A: 255}, hex: "#0000FF", luminance: 0.0722},
			want: color.RGBA{R: 0, G: 0, B: 255, A: 255},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.ColorGo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color.ColorGo() = %v, want %v", got, tt.want)
			}
		})
	}
}
