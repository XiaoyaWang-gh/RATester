package captcha

import (
	"fmt"
	"testing"
)

func TestMin3_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	x := uint8(1)
	y := uint8(2)
	z := uint8(3)
	m := min3(x, y, z)
	if m != 1 {
		t.Errorf("min3(%d, %d, %d) = %d, want %d", x, y, z, m, 1)
	}
}
