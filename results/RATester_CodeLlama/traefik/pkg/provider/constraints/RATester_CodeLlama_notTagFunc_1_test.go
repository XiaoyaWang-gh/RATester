package constraints

import (
	"fmt"
	"testing"
)

func TestNotTagFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := func(tags []string) bool {
		return true
	}
	b := notTagFunc(a)
	if b([]string{"a", "b"}) {
		t.Error("notTagFunc failed")
	}
}
