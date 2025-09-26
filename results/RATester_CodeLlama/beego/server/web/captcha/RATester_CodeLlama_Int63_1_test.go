package captcha

import (
	"fmt"
	"testing"
)

func TestInt63_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &siprng{}
	p.rekey()
	for i := 0; i < 1000; i++ {
		x := p.Int63()
		if x < 0 {
			t.Errorf("Int63() = %d; want >= 0", x)
		}
	}
}
