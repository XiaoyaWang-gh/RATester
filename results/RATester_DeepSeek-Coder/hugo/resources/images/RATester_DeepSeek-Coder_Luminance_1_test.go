package images

import (
	"fmt"
	"image/color"
	"testing"
)

func TestLuminance_1(t *testing.T) {
	tests := []struct {
		name string
		c    Color
		want float64
	}{
		{
			name: "Test black",
			c:    Color{color: color.Black, hex: "#000000", luminance: 0.0},
			want: 0.0,
		},
		{
			name: "Test white",
			c:    Color{color: color.White, hex: "#FFFFFF", luminance: 1.0},
			want: 1.0,
		},
		{
			name: "Test red",
			c:    Color{color: color.RGBA{R: 255, A: 255}, hex: "#FF0000", luminance: 0.2126},
			want: 0.2126,
		},
		{
			name: "Test green",
			c:    Color{color: color.RGBA{G: 255, A: 255}, hex: "#00FF00", luminance: 0.7152},
			want: 0.7152,
		},
		{
			name: "Test blue",
			c:    Color{color: color.RGBA{B: 255, A: 255}, hex: "#0000FF", luminance: 0.0722},
			want: 0.0722,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.Luminance(); got != tt.want {
				t.Errorf("Luminance() = %v, want %v", got, tt.want)
			}
		})
	}
}
