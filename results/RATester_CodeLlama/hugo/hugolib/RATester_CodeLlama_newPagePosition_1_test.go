package hugolib

import (
	"fmt"
	"testing"
)

func TestNewPagePosition_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &nextPrev{}
	p := newPagePosition(n)
	if p.nextPrev != n {
		t.Errorf("p.nextPrev != n")
	}
}
