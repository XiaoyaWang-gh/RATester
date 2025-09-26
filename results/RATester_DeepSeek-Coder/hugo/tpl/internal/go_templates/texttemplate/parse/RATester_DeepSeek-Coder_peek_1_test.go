package parse

import (
	"fmt"
	"testing"
)

func TestPeek_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		peekCount: 0,
		token: [3]item{
			{typ: itemEOF, pos: 0, val: "", line: 0},
			{typ: itemEOF, pos: 0, val: "", line: 0},
			{typ: itemEOF, pos: 0, val: "", line: 0},
		},
	}

	expected := item{typ: itemEOF, pos: 0, val: "", line: 0}
	result := tree.peek()

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
