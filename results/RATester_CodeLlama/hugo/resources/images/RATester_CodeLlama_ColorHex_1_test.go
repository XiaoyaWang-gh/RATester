package images

import (
	"fmt"
	"testing"
)

func TestColorHex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := Color{hex: "#000000"}
	if got := c.ColorHex(); got != "#000000" {
		t.Errorf("ColorHex() = %v, want #000000", got)
	}
}
