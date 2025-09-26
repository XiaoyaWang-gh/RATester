package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCopyList_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &ListNode{
		Nodes: []Node{
			&TextNode{
				Text: []byte("foo"),
			},
			&TextNode{
				Text: []byte("bar"),
			},
		},
	}
	l2 := l.CopyList()
	if !reflect.DeepEqual(l, l2) {
		t.Errorf("l and l2 are not equal")
	}
	if l == l2 {
		t.Errorf("l and l2 are the same")
	}
	if l.Nodes[0] == l2.Nodes[0] {
		t.Errorf("l.Nodes[0] and l2.Nodes[0] are the same")
	}
	if l.Nodes[1] == l2.Nodes[1] {
		t.Errorf("l.Nodes[1] and l2.Nodes[1] are the same")
	}
}
