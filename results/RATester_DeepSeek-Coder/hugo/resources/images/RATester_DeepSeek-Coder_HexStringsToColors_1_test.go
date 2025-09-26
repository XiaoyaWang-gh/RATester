package images

import (
	"fmt"
	"testing"
)

func TestHexStringsToColors_1(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []Color
	}{
		{
			name: "Test with valid hex strings",
			args: []string{"#FFFFFF", "#000000", "#FF0000"},
			want: []Color{
				{hex: "#FFFFFF", luminance: 1},
				{hex: "#000000", luminance: 0},
				{hex: "#FF0000", luminance: 0.2126},
			},
		},
		{
			name: "Test with invalid hex strings",
			args: []string{"#FFFFFG", "#00000Z", "#12345"},
			want: []Color{
				{hex: "#000000", luminance: 0},
				{hex: "#000000", luminance: 0},
				{hex: "#000000", luminance: 0},
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

			got := HexStringsToColors(tt.args...)
			for i, g := range got {
				if g.ColorHex() != tt.want[i].ColorHex() || g.Luminance() != tt.want[i].Luminance() {
					t.Errorf("HexStringsToColors() = %v, want %v", g, tt.want[i])
				}
			}
		})
	}
}
