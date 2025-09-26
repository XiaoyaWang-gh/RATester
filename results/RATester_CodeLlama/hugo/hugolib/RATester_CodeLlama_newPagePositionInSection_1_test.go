package hugolib

import (
	"fmt"
	"testing"
)

func TestNewPagePositionInSection_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &nextPrev{}
	ppis := newPagePositionInSection(n)
	if ppis.nextPrev != n {
		t.Errorf("ppis.nextPrev != n")
	}
}
