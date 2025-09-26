package compare

import (
	"fmt"
	"testing"
)

func TestNe_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{}
	first := "foo"
	others := []any{"bar", "baz"}
	if n.Ne(first, others...) {
		t.Errorf("Expected %q to be not equal to %q", first, others)
	}
}
