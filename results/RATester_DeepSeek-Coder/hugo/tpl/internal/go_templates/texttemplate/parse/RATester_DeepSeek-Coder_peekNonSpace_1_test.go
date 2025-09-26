package parse

import (
	"fmt"
	"testing"
)

func TestPeekNonSpace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		Name: "test",
		lex: &lexer{
			input: "test input",
		},
	}

	expected := item{
		typ: itemText,
		val: "test",
	}

	result := tree.peekNonSpace()

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
