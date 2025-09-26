package template

import (
	"fmt"
	"testing"
)

func TestEatWhiteSpace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := []byte("  \t\n\f\r")
	i := 0
	if eatWhiteSpace(s, i) != len(s) {
		t.Errorf("eatWhiteSpace(%v, %v) = %v; want %v", s, i, eatWhiteSpace(s, i), len(s))
	}
}
