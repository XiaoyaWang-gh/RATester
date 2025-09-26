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

	l := &pageLexer{
		input: []byte("test input"),
		pos:   5,
		start: 3,
	}

	l.ignore()

	if l.start != 5 {
		t.Errorf("Expected start to be 5, got %d", l.start)
	}
}
