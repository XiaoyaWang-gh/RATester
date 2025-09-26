package csrf

import (
	"fmt"
	"testing"
)

func TestCompareStrings_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := "a"
	b := "b"
	if compareStrings(a, b) {
		t.Errorf("compareStrings(%s, %s) = true, want false", a, b)
	}
	a = "a"
	b = "a"
	if !compareStrings(a, b) {
		t.Errorf("compareStrings(%s, %s) = false, want true", a, b)
	}
}
