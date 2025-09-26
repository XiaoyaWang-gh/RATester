package captcha

import (
	"fmt"
	"testing"
)

func TestIntn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &siprng{}
	p.rekey()
	for i := 0; i < 1000; i++ {
		n := p.Intn(1000)
		if n < 0 || n >= 1000 {
			t.Errorf("Intn(%d) = %d", 1000, n)
		}
	}
}
