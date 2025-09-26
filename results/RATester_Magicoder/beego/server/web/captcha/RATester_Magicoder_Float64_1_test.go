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

	rng := &siprng{}
	rng.rekey()

	min := float64(0)
	max := float64(1)

	for i := 0; i < 1000; i++ {
		result := rng.Float64()
		if result < min || result > max {
			t.Errorf("Float64() = %f, want in range [%f, %f]", result, min, max)
		}
	}
}
