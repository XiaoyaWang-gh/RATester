package parse

import (
	"fmt"
	"testing"
)

func TestThisItem_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &lexer{
		name:      "test",
		input:     "test input",
		start:     0,
		pos:       5,
		startLine: 1,
		line:      1,
	}

	expected := item{
		typ:  itemType(1),
		pos:  0,
		val:  "test",
		line: 1,
	}

	result := l.thisItem(itemType(1))

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	if l.start != 5 {
		t.Errorf("Expected start to be 5, got %d", l.start)
	}

	if l.startLine != 1 {
		t.Errorf("Expected startLine to be 1, got %d", l.startLine)
	}
}
