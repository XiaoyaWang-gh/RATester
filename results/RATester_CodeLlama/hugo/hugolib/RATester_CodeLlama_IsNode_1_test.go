package hugolib

import (
	"fmt"
	"testing"
)

func TestIsNode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	if p.IsNode() {
		t.Errorf("IsNode() = true, want false")
	}
}
