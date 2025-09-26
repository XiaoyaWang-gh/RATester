package constraints

import (
	"fmt"
	"testing"
)

func TestAndTagFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := func(tags []string) bool {
		return true
	}
	b := func(tags []string) bool {
		return true
	}
	c := andTagFunc(a, b)
	if !c([]string{"a", "b"}) {
		t.Errorf("andTagFunc failed")
	}
}
