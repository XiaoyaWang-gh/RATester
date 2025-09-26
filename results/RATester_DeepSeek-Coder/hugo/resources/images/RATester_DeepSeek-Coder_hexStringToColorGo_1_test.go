package images

import (
	"fmt"
	"image/color"
	"testing"
)

func TestHexStringToColorGo_1(t *testing.T) {
	t.Run("valid hex string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		tests := []struct {
			input    string
			expected color.Color
		}{
			{"#ffffff", color.White},
			{"#000000", color.Black},
			{"#abc", color.RGBA{0xAA, 0xBB, 0xCC, 0xFF}},
			{"#abcd", color.RGBA{0xAA, 0xBB, 0xCC, 0xDD}},
			{"#abcdef", color.RGBA{0xAB, 0xCD, 0xEF, 0xFF}},
			{"#abcdef00", color.RGBA{0xAB, 0xCD, 0xEF, 0x00}},
		}

		for _, test := range tests {
			result, err := hexStringToColorGo(test.input)
			if err != nil {
				t.Errorf("hexStringToColorGo(%q) returned error: %v", test.input, err)
				continue
			}

			if result != test.expected {
				t.Errorf("hexStringToColorGo(%q) = %v, want %v", test.input, result, test.expected)
			}
		}
	})

	t.Run("invalid hex string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		tests := []string{
			"#abcg",
			"#abcdf",
			"#abcdefg",
			"#12345",
			"#1234567",
		}

		for _, test := range tests {
			_, err := hexStringToColorGo(test)
			if err == nil {
				t.Errorf("hexStringToColorGo(%q) did not return error", test)
			}
		}
	})
}
