package images

import (
	"fmt"
	"testing"
)

func TestColorHex_1(t *testing.T) {
	tests := []struct {
		name string
		c    Color
		want string
	}{
		{
			name: "Test white color",
			c:    Color{hex: "#FFFFFF"},
			want: "#FFFFFF",
		},
		{
			name: "Test black color",
			c:    Color{hex: "#000000"},
			want: "#000000",
		},
		{
			name: "Test red color",
			c:    Color{hex: "#FF0000"},
			want: "#FF0000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.ColorHex(); got != tt.want {
				t.Errorf("Color.ColorHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
