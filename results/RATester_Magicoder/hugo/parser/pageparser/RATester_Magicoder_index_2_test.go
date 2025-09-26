package pageparser

import (
	"fmt"
	"testing"
)

func Testindex_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{
		input: []byte("hello world"),
	}
	sep := []byte("world")
	expected := 6
	result := l.index(sep)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
