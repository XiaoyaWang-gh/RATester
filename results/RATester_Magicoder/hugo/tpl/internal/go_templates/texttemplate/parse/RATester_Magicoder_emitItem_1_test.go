package parse

import (
	"fmt"
	"testing"
)

func TestemitItem_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &lexer{
		name: "test",
	}
	i := item{
		typ:  itemEOF,
		pos:  0,
		val:  "EOF",
		line: 1,
	}
	expected := i
	result := l.emitItem(i)
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
	if l.item != expected {
		t.Errorf("Expected %v, got %v", expected, l.item)
	}
}
