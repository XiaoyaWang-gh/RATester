package captcha

import (
	"fmt"
	"testing"
)

func TestInt63n_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &siprng{}
	p.rekey()
	for i := 0; i < 1000; i++ {
		n := int64(i + 1)
		v := p.Int63n(n)
		if v < 0 || v >= n {
			t.Errorf("Int63n(%d) = %d", n, v)
		}
	}
}
