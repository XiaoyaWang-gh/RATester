package captcha

import (
	"fmt"
	"testing"
)

func TestIntn_1(t *testing.T) {
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
		rng.Intn(0)
	})

	t.Run("n > 0", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		n := 100
		for i := 0; i < 1000; i++ {
			result := rng.Intn(n)
			if result < 0 || result >= n {
				t.Errorf("Intn(%d) = %d; want 0 <= result < %d", n, result, n)
			}
		}
	})
}
