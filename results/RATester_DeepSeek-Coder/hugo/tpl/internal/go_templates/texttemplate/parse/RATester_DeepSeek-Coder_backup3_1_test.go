package parse

import (
	"fmt"
	"testing"
)

func TestBackup3_1(t *testing.T) {
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

	t2 := item{typ: itemNumber, val: "4"}
	t1 := item{typ: itemNumber, val: "5"}

	tree.backup3(t2, t1)

	if tree.token[0].val != "5" || tree.token[1].val != "4" || tree.token[2].val != "3" {
		t.Errorf("Expected tokens to be [5, 4, 3], got %v", tree.token)
	}
	if tree.peekCount != 3 {
		t.Errorf("Expected peekCount to be 3, got %d", tree.peekCount)
	}
}
