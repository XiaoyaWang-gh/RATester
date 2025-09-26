package pageparser

import (
	"fmt"
	"testing"
)

func Testignore_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{
		input: []byte("test input"),
		start: 0,
	}
	l.ignore()
	if l.start != 0 {
		t.Errorf("Expected start to be 0, but got %d", l.start)
	}
}
