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

	rng := &siprng{
		k0:  0x0102030405060708,
		k1:  0x090a0b0c0d0e0f10,
		ctr: 0x1122334455667788,
	}

	expected := uint32(rng.Uint64())
	actual := rng.Uint32()

	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
