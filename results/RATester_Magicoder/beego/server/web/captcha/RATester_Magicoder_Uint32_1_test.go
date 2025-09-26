package captcha

import (
	"fmt"
	"testing"
)

func TestUint32_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rng := &siprng{}
	rng.k0 = 1
	rng.k1 = 2
	rng.ctr = 3

	expected := uint32(rng.Uint64())
	result := rng.Uint32()

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
