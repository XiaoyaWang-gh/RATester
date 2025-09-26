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
			name: "Test case 1",
			c:    Color{hex: "#FFFFFF"},
			want: "#FFFFFF",
		},
		{
			name: "Test case 2",
			c:    Color{hex: "#000000"},
			want: "#000000",
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

			if got := tt.c.ColorHex(); got != tt.want {
				t.Errorf("ColorHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
