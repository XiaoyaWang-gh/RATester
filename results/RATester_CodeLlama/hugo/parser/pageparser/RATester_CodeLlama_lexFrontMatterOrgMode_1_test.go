package pageparser

import (
	"fmt"
	"testing"
)

func TestLexFrontMatterOrgMode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := new(pageLexer)
	l.input = []byte("+++\n+++\n")
	l.run()
	if len(l.items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(l.items))
	}
	if l.items[0].Type != TypeFrontMatterORG {
		t.Errorf("Expected TypeFrontMatterORG, got %s", l.items[0].Type)
	}
	if l.items[1].Type != TypeFrontMatterORG {
		t.Errorf("Expected TypeFrontMatterORG, got %s", l.items[1].Type)
	}
}
