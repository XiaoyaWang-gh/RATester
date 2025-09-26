package pageparser

import (
	"bytes"
	"fmt"
	"testing"
)

func TestInput_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{
		input: []byte("test input"),
	}

	expected := []byte("test input")
	result := l.Input()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
