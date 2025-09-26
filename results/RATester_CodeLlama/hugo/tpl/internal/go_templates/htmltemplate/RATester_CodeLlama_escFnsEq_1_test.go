package template

import (
	"fmt"
	"testing"
)

func TestEscFnsEq_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := "a"
	b := "a"
	if !escFnsEq(a, b) {
		t.Errorf("a and b should be equal")
	}
}
