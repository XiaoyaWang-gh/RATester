package captcha

import (
	"fmt"
	"testing"
)

func TestUint64_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rng := &siprng{
		k0:  0x0706050403020100,
		k1:  0x0f0e0d0c0b0a0908,
		ctr: 0,
	}

	expected := uint64(0x0706050403020100)
	actual := rng.Uint64()
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
