package images

import (
	"fmt"
	"image/color"
	"reflect"
	"testing"
)

func TestHexStringsToColors_1(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []Color
	}{
		{
			name: "Test case 1",
			args: []string{"#FF0000", "#00FF00", "#0000FF"},
			want: []Color{
				{color: color.RGBA{R: 255, G: 0, B: 0, A: 255}, hex: "#FF0000", luminance: 0.2126},
				{color: color.RGBA{R: 0, G: 255, B: 0, A: 255}, hex: "#00FF00", luminance: 0.7152},
				{color: color.RGBA{R: 0, G: 0, B: 255, A: 255}, hex: "#0000FF", luminance: 0.0722},
			},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := HexStringsToColors(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexStringsToColors() = %v, want %v", got, tt.want)
			}
		})
	}
}
