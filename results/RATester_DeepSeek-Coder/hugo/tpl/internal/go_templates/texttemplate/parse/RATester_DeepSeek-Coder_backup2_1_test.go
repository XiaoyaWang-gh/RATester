package parse

import (
	"fmt"
	"testing"
)

func TestBackup2_1(t *testing.T) {
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
		peekCount: 3,
	}

	tree.backup2(item{typ: itemNumber, val: "4"})

	if tree.peekCount != 2 {
		t.Errorf("Expected peekCount to be 2, got %d", tree.peekCount)
	}

	if tree.token[0].val != "4" {
		t.Errorf("Expected token[0].val to be '4', got '%s'", tree.token[0].val)
	}

	if tree.token[1].val != "1" {
		t.Errorf("Expected token[1].val to be '1', got '%s'", tree.token[1].val)
	}

	if tree.token[2].val != "2" {
		t.Errorf("Expected token[2].val to be '2', got '%s'", tree.token[2].val)
	}
}
