package parse

import (
	"fmt"
	"testing"
)

func TestpopVars_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		vars: []string{"var1", "var2", "var3"},
	}
	tree.popVars(2)
	if len(tree.vars) != 1 || tree.vars[0] != "var3" {
		t.Errorf("Expected vars to be [var3], but got %v", tree.vars)
	}
}
