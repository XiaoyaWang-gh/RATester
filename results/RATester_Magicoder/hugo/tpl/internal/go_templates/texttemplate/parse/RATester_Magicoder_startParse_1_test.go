package parse

import (
	"fmt"
	"testing"
)

func TeststartParse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	lex := &lexer{}
	funcs := []map[string]any{}
	treeSet := map[string]*Tree{}

	tree.startParse(funcs, lex, treeSet)

	if tree.Root != nil {
		t.Errorf("Expected Root to be nil, but got %v", tree.Root)
	}
	if tree.lex != lex {
		t.Errorf("Expected lex to be %v, but got %v", lex, tree.lex)
	}
	if len(tree.vars) != 1 || tree.vars[0] != "$" {
		t.Errorf("Expected vars to be [\"$\"], but got %v", tree.vars)
	}
	if len(tree.funcs) != 0 {
		t.Errorf("Expected funcs to be empty, but got %v", tree.funcs)
	}
	if len(tree.treeSet) != 0 {
		t.Errorf("Expected treeSet to be empty, but got %v", tree.treeSet)
	}
	if tree.lex.options.emitComment != (tree.Mode&ParseComments != 0) {
		t.Errorf("Expected emitComment to be %v, but got %v", tree.Mode&ParseComments != 0, tree.lex.options.emitComment)
	}
	if tree.lex.options.breakOK != !tree.hasFunction("break") {
		t.Errorf("Expected breakOK to be %v, but got %v", !tree.hasFunction("break"), tree.lex.options.breakOK)
	}
	if tree.lex.options.continueOK != !tree.hasFunction("continue") {
		t.Errorf("Expected continueOK to be %v, but got %v", !tree.hasFunction("continue"), tree.lex.options.continueOK)
	}
}
