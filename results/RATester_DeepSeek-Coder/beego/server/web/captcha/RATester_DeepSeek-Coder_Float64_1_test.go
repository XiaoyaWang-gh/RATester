package captcha

import (
	"fmt"
	"testing"
)

func TestFloat64_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	rng := &siprng{}

	for i := 0; i < 100; i++ {
		f := rng.Float64()

		if f < 0 || f >= 1 {
			t.Errorf("rng.Float64() = %f; want 0 <= f < 1", f)
		}
	}
}
