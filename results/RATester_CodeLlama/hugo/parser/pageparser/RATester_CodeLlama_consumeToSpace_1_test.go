package pageparser

import (
	"fmt"
	"testing"
)

func TestConsumeToSpace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{}
	l.input = []byte("abc")
	l.pos = 0
	l.consumeToSpace()
	if l.pos != 3 {
		t.Errorf("l.pos = %d, want 3", l.pos)
	}
}
