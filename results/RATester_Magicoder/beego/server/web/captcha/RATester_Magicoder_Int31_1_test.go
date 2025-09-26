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

	rng := &siprng{}
	rng.k0 = 1
	rng.k1 = 2
	rng.ctr = 3

	result := rng.Int31()

	if result < 0 || result > 0x7fffffff {
		t.Errorf("Int31() returned an invalid value: %d", result)
	}
}
