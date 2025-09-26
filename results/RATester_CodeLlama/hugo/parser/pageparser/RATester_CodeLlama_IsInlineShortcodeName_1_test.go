package pageparser

import (
	"fmt"
	"testing"
)

func TestIsInlineShortcodeName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := Item{Type: tScNameInline}
	if !i.IsInlineShortcodeName() {
		t.Errorf("IsInlineShortcodeName failed")
	}
}
