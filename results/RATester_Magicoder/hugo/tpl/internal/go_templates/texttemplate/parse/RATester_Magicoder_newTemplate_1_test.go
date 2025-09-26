package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewTemplate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	pos := Pos(1)
	line := 1
	name := "test"
	pipe := &PipeNode{}
	expected := &TemplateNode{tr: tree, NodeType: NodeTemplate, Pos: pos, Line: line, Name: name, Pipe: pipe}
	result := tree.newTemplate(pos, line, name, pipe)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
