package parse

import (
	"fmt"
	"testing"
)

func Testbackup3_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	t2 := item{typ: 1, val: "t2"}
	t1 := item{typ: 2, val: "t1"}
	tree.backup3(t2, t1)

	if tree.token[1].val != "t1" {
		t.Errorf("Expected token[1].val to be 't1', but got %s", tree.token[1].val)
	}
	if tree.token[2].val != "t2" {
		t.Errorf("Expected token[2].val to be 't2', but got %s", tree.token[2].val)
	}
	if tree.peekCount != 3 {
		t.Errorf("Expected peekCount to be 3, but got %d", tree.peekCount)
	}
}
