package compare

import (
	"fmt"
	"testing"
)

func TestLt_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{}
	first := "a"
	others := []any{"b", "c"}
	if !n.Lt(first, others...) {
		t.Errorf("Lt failed")
	}
}
