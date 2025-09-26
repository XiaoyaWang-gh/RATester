package pageparser

import (
	"fmt"
	"testing"
)

func TestConsumeSpace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{}
	l.input = []byte("  ")
	l.consumeSpace()
	if l.pos != 2 {
		t.Errorf("l.pos = %d, want 2", l.pos)
	}
}
