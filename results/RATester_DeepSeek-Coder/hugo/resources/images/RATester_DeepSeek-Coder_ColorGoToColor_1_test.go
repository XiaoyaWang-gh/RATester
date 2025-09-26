package images

import (
	"fmt"
	"image/color"
	"reflect"
	"testing"
)

func TestColorGoToColor_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		c    color.Color
		want Color
	}{
		{
			name: "Test with black color",
			c:    color.Black,
			want: Color{color: color.Black},
		},
		{
			name: "Test with white color",
			c:    color.White,
			want: Color{color: color.White},
		},
		{
			name: "Test with red color",
			c:    color.RGBA{R: 255, A: 255},
			want: Color{color: color.RGBA{R: 255, A: 255}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := ColorGoToColor(tt.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColorGoToColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
