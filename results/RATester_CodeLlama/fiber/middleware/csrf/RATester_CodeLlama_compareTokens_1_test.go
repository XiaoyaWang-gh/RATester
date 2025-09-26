package csrf

import (
	"fmt"
	"testing"
)

func TestCompareTokens_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := []byte("a")
	b := []byte("b")
	if compareTokens(a, b) {
		t.Errorf("compareTokens(%v, %v) = true, want false", a, b)
	}
	if !compareTokens(a, a) {
		t.Errorf("compareTokens(%v, %v) = false, want true", a, a)
	}
}
