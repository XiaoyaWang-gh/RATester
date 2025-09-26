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

	rng := &siprng{}
	n := int32(10)

	// Test for panic when n <= 0
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic when n <= 0")
		}
	}()
	rng.Int31n(0)

	// Test for valid input
	for i := 0; i < 1000; i++ {
		result := rng.Int31n(n)
		if result < 0 || result >= n {
			t.Errorf("Int31n(%d) = %d, want: 0 <= result < %d", n, result, n)
		}
	}
}
