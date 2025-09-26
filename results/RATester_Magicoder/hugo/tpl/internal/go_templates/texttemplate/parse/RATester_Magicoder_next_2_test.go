package parse

import (
	"fmt"
	"testing"
)

func Testnext_2(t *testing.T) {
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
		peekCount: 0,
	}

	expected := item{typ: itemNumber, val: "1"}
	result := tree.next()

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	if tree.peekCount != 2 {
		t.Errorf("Expected peekCount to be 2, but got %d", tree.peekCount)
	}
}
