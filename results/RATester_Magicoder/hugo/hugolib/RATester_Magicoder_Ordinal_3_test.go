package hugolib

import (
	"fmt"
	"testing"
)

func TestOrdinal_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	page := pageWithOrdinal{ordinal: 1}

	if page.Ordinal() != 1 {
		t.Errorf("Expected ordinal to be 1, but got %d", page.Ordinal())
	}
}
