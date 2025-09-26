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
		for i := 0; i < 100; i++ {
			v := rng.Int63n(n)
			if v < 0 || v >= n {
				t.Errorf("Int63n(%d) = %d; want 0 <= v < %d", n, v, n)
			}
		}
	})
}
