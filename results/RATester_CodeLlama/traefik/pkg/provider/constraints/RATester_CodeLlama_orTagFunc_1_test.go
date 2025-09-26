package constraints

import (
	"fmt"
	"testing"
)

func TestOrTagFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := func(tags []string) bool {
		return true
	}
	b := func(tags []string) bool {
		return false
	}
	c := orTagFunc(a, b)
	if !c([]string{"a", "b"}) {
		t.Error("expected true")
	}
	if c([]string{"c", "d"}) {
		t.Error("expected false")
	}
}
