package pageparser

import (
	"fmt"
	"testing"
)

func TestIsEOF_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{}
	l.input = []byte("")
	l.pos = 0
	if l.isEOF() {
		t.Errorf("Expected isEOF to return false")
	}
	l.input = []byte("a")
	l.pos = 1
	if !l.isEOF() {
		t.Errorf("Expected isEOF to return true")
	}
}
