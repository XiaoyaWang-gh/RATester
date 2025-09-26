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

	p := pageWithOrdinal{
		ordinal: 1,
	}

	if p.Ordinal() != 1 {
		t.Errorf("expected 1, got %d", p.Ordinal())
	}
}
