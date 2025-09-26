package gin

import (
	"fmt"
	"testing"
)

func TestAddChild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &node{}
	child := &node{}
	n.addChild(child)
	if len(n.children) != 1 {
		t.Errorf("TestaddChild failed")
	}
}
