package parse

import (
	"fmt"
	"testing"
)

func TestPopVars_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		vars: []string{"var1", "var2", "var3"},
	}

	tree.popVars(2)

	if len(tree.vars) != 1 {
		t.Errorf("Expected length of vars to be 1, got %d", len(tree.vars))
	}

	if tree.vars[0] != "var3" {
		t.Errorf("Expected last variable to be 'var3', got %s", tree.vars[0])
	}
}
