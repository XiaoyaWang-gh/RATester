package pageparser

import (
	"fmt"
	"testing"
)

func TestIndex_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{}
	l.input = []byte("abc")
	l.pos = 1
	sep := []byte("c")
	if l.index(sep) != 2 {
		t.Errorf("l.index(sep) = %d, want %d", l.index(sep), 2)
	}
}
