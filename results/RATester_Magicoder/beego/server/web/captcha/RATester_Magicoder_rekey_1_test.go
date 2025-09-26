package captcha

import (
	"fmt"
	"testing"
)

func Testrekey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rng := &siprng{}
	rng.rekey()

	if rng.k0 == 0 || rng.k1 == 0 {
		t.Errorf("rekey() did not generate valid keys")
	}
}
