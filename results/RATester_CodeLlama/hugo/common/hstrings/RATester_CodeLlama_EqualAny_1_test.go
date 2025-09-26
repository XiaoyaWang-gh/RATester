package hstrings

import (
	"fmt"
	"testing"
)

func TestEqualAny_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := "a"
	b := []string{"a", "b", "c"}
	if !EqualAny(a, b...) {
		t.Errorf("Expected true, got false")
	}
}
