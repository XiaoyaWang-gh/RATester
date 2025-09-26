package captcha

import (
	"fmt"
	"testing"
)

func TestInt63_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rng := &siprng{}
	rng.k0 = 1
	rng.k1 = 2
	rng.ctr = 3

	result := rng.Int63()

	if result < 0 || result > 9223372036854775807 {
		t.Errorf("Int63() returned an invalid value: %d", result)
	}
}
