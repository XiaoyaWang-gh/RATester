package parse

import (
	"fmt"
	"testing"
)

func TestNewPipeline_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	vars := []*VariableNode{
		{Ident: []string{"var1"}},
		{Ident: []string{"var2"}},
	}
	pos := Pos(1)
	line := 2

	pipeline := tree.newPipeline(pos, line, vars)

	if pipeline.tr != tree {
		t.Errorf("Expected tree to be %v, got %v", tree, pipeline.tr)
	}
	if pipeline.NodeType != NodePipe {
		t.Errorf("Expected NodeType to be %v, got %v", NodePipe, pipeline.NodeType)
	}
	if pipeline.Pos != pos {
		t.Errorf("Expected Pos to be %v, got %v", pos, pipeline.Pos)
	}
	if pipeline.Line != line {
		t.Errorf("Expected Line to be %v, got %v", line, pipeline.Line)
	}
	if len(pipeline.Decl) != len(vars) {
		t.Errorf("Expected Decl to have length %v, got %v", len(vars), len(pipeline.Decl))
	}
	for i, v := range vars {
		if pipeline.Decl[i] != v {
			t.Errorf("Expected Decl[%v] to be %v, got %v", i, v, pipeline.Decl[i])
		}
	}
}
