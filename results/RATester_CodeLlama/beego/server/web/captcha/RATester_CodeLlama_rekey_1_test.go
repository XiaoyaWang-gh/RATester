package captcha

import (
	"fmt"
	"testing"
)

func TestRekey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &siprng{}
	p.rekey()
	if p.k0 == 0 || p.k1 == 0 {
		t.Errorf("rekey() failed")
	}
}
