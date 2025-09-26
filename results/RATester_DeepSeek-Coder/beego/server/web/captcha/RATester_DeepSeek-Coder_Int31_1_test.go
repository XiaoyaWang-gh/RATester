package captcha

import (
	"fmt"
	"testing"
)

func TestInt31_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rng := &siprng{
		k0:  0x0706050403020100,
		k1:  0x0f0e0d0c0b0a0908,
		ctr: 0x0706050403020100,
	}

	expected := int32(0x7fffffff)
	result := rng.Int31()

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
