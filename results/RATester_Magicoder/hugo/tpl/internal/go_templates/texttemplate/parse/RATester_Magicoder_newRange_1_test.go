package parse

import (
	"fmt"
	"testing"
)

func TestnewRange_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	pos := Pos(1)
	line := 1
	pipe := &PipeNode{}
	list := &ListNode{}
	elseList := &ListNode{}

	result := tree.newRange(pos, line, pipe, list, elseList)

	if result.Pos != pos {
		t.Errorf("Expected Pos to be %v, but got %v", pos, result.Pos)
	}

	if result.Line != line {
		t.Errorf("Expected Line to be %v, but got %v", line, result.Line)
	}

	if result.Pipe != pipe {
		t.Errorf("Expected Pipe to be %v, but got %v", pipe, result.Pipe)
	}

	if result.List != list {
		t.Errorf("Expected List to be %v, but got %v", list, result.List)
	}

	if result.ElseList != elseList {
		t.Errorf("Expected ElseList to be %v, but got %v", elseList, result.ElseList)
	}
}
