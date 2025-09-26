package parse

import (
	"fmt"
	"testing"
)

func TestNext_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		peekCount: 2,
		token: [3]item{
			{typ: itemNumber, val: "1"},
			{typ: itemNumber, val: "2"},
			{typ: itemNumber, val: "3"},
		},
	}

	expected := item{typ: itemNumber, val: "1"}
	result := tree.next()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	expected = item{typ: itemNumber, val: "2"}
	result = tree.next()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	expected = item{typ: itemNumber, val: "3"}
	result = tree.next()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
