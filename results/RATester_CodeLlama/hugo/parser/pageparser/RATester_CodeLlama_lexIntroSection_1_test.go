package pageparser

import (
	"fmt"
	"testing"
)

func TestLexIntroSection_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{
		input: []byte("+++\n+++\n"),
	}
	l.run()
	if l.items[0].Type != TypeFrontMatterTOML {
		t.Errorf("Expected TypeFrontMatterTOML, got %v", l.items[0].Type)
	}
	if l.items[1].Type != TypeFrontMatterTOML {
		t.Errorf("Expected TypeFrontMatterTOML, got %v", l.items[1].Type)
	}
}
