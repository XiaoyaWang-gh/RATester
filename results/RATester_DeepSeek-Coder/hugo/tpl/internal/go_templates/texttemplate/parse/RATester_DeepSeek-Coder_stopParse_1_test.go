package parse

import (
	"fmt"
	"testing"
)

func TestStopParse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		lex:     &lexer{},
		vars:    []string{"var1", "var2"},
		funcs:   []map[string]any{{"func1": func() {}}},
		treeSet: map[string]*Tree{"tree1": {}},
	}

	tree.stopParse()

	if tree.lex != nil {
		t.Errorf("Expected lex to be nil, got %v", tree.lex)
	}
	if len(tree.vars) != 0 {
		t.Errorf("Expected vars to be empty, got %v", tree.vars)
	}
	if len(tree.funcs) != 0 {
		t.Errorf("Expected funcs to be empty, got %v", tree.funcs)
	}
	if tree.treeSet != nil {
		t.Errorf("Expected treeSet to be nil, got %v", tree.treeSet)
	}
}
