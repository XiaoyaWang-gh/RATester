package tcp

import (
	"fmt"
	"testing"
)

func TestGcd_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := 10
	b := 15
	if gcd(a, b) != 5 {
		t.Errorf("gcd(%d, %d) = %d, want %d", a, b, gcd(a, b), 5)
	}
}
