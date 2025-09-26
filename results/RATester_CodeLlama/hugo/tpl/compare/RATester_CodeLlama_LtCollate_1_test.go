package compare

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/langs"
)

func TestLtCollate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{}
	collator := &langs.Collator{}
	first := "a"
	others := []any{"b", "c"}
	if !n.LtCollate(collator, first, others...) {
		t.Errorf("LtCollate failed")
	}
}
