package parse

import (
	"fmt"
	"testing"
)

func Testbackup2_1(t *testing.T) {
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

	expected := [3]item{
		{typ: itemNumber, val: "1"},
		{typ: itemNumber, val: "2"},
		{typ: itemNumber, val: "3"},
	}

	tree.backup2(item{typ: itemNumber, val: "4"})

	if tree.token[0].val != expected[0].val || tree.token[1].val != expected[1].val || tree.token[2].val != expected[2].val {
		t.Errorf("Expected token[0] to be %s, got %s", expected[0].val, tree.token[0].val)
		t.Errorf("Expected token[1] to be %s, got %s", expected[1].val, tree.token[1].val)
		t.Errorf("Expected token[2] to be %s, got %s", expected[2].val, tree.token[2].val)
	}

	if tree.peekCount != 2 {
		t.Errorf("Expected peekCount to be %d, got %d", 2, tree.peekCount)
	}
}
