package captcha

import (
	"fmt"
	"testing"
)

func TestInt63n_1(t *testing.T) {
	rng := &siprng{}

	t.Run("n <= 0", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		rng.Int63n(0)
	})

	t.Run("n > 0", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		n := int64(10)
		for i := 0; i < 1000; i++ {
			result := rng.Int63n(n)
			if result < 0 || result >= n {
				t.Errorf("Int63n(%d) = %d, want a value in [0, %d)", n, result, n)
			}
		}
	})
}
