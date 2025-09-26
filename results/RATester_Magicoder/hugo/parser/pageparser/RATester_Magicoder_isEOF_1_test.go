package pageparser

import (
	"fmt"
	"testing"
)

func TestisEOF_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{
		input: []byte("test"),
		pos:   4,
	}

	if !l.isEOF() {
		t.Errorf("Expected isEOF to return true, but got false")
	}

	l.pos = 3
	if l.isEOF() {
		t.Errorf("Expected isEOF to return false, but got true")
	}
}
