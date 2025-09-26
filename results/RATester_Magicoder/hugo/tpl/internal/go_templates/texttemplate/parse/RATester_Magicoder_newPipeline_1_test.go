package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewPipeline_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	pos := Pos(1)
	line := 1
	vars := []*VariableNode{
		{
			Ident: []string{"var1"},
		},
		{
			Ident: []string{"var2"},
		},
	}
	expected := &PipeNode{
		tr:       tree,
		NodeType: NodePipe,
		Pos:      pos,
		Line:     line,
		Decl:     vars,
	}
	result := tree.newPipeline(pos, line, vars)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
