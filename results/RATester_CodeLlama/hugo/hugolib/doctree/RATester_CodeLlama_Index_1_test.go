package doctree

import (
	"fmt"
	"testing"
)

func TestIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := DimensionFlag(1)
	if d.Index() != 0 {
		t.Errorf("expected 0, got %d", d.Index())
	}
}
