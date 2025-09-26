package pageparser

import (
	"fmt"
	"testing"
)

func TestIgnore_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{}
	l.ignore()
	if l.start != l.pos {
		t.Errorf("l.start = %d, l.pos = %d", l.start, l.pos)
	}
}
