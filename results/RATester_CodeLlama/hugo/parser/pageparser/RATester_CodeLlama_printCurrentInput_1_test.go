package pageparser

import (
	"fmt"
	"testing"
)

func TestPrintCurrentInput_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{}
	l.input = []byte("input")
	l.pos = 1
	l.printCurrentInput()
	if l.pos != 1 {
		t.Errorf("l.pos = %d, want 1", l.pos)
	}
}
