package parse

import (
	"fmt"
	"testing"
)

func TestNextNonSpace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		lex: &lexer{
			input: "  hello world",
		},
	}

	token := tree.nextNonSpace()
	if token.typ != itemSpace {
		t.Errorf("Expected itemSpace, got %v", token.typ)
	}
	if token.val != "  " {
		t.Errorf("Expected '  ', got '%v'", token.val)
	}

	token = tree.nextNonSpace()
	if token.typ != itemText {
		t.Errorf("Expected itemText, got %v", token.typ)
	}
	if token.val != "hello" {
		t.Errorf("Expected 'hello', got '%v'", token.val)
	}
}
