package pageparser

import (
	"fmt"
	"testing"
)

func TestHasPrefix_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{}
	l.input = []byte("abc")
	l.pos = 1
	prefix := []byte("ab")
	if !l.hasPrefix(prefix) {
		t.Errorf("Expected true, got false")
	}
}
