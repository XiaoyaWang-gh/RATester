package parse

import (
	"fmt"
	"testing"
)

func TestParseDefinition_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		Name: "test",
		lex: &lexer{
			input: "test",
		},
	}
	tree.parseDefinition()
	if tree.Name != "test" {
		t.Errorf("Expected Name to be 'test', got %s", tree.Name)
	}
}
