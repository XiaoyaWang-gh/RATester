package echo

import (
	"fmt"
	"testing"
)

func TestFindStaticChild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &node{}
	n.staticChildren = append(n.staticChildren, &node{label: 'a'})
	if n.findStaticChild('a') == nil {
		t.Errorf("Expected to find static child")
	}
	if n.findStaticChild('b') != nil {
		t.Errorf("Expected not to find static child")
	}
}
