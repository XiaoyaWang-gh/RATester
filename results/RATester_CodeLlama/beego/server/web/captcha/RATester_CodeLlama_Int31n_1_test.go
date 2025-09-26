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

	p := &siprng{}
	p.rekey()
	for i := 0; i < 1000; i++ {
		n := p.Int31n(100)
		if n < 0 || n >= 100 {
			t.Errorf("n = %d; want [0, 100)", n)
		}
	}
}
