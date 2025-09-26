package parse

import (
	"fmt"
	"testing"
)

func Testpeek_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		token: [3]item{
			{typ: itemNumber, val: "1"},
			{typ: itemNumber, val: "2"},
			{typ: itemNumber, val: "3"},
		},
		peekCount: 2,
	}

	expected := item{typ: itemNumber, val: "2"}
	result := tree.peek()

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	if tree.peekCount != 3 {
		t.Errorf("Expected peekCount to be 3, but got %d", tree.peekCount)
	}
}
