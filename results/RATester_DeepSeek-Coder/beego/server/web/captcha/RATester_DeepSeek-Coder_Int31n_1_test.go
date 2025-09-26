package captcha

import (
	"fmt"
	"testing"
)

func TestInt31n_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rng := &siprng{
		k0:  1234567890,
		k1:  362436069,
		ctr: 521288629,
	}

	n := int32(100)
	result := rng.Int31n(n)

	if result < 0 || result >= n {
		t.Errorf("Int31n(%d) = %d; want 0 <= result < %d", n, result, n)
	}
}
