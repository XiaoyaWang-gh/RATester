package parse

import (
	"fmt"
	"testing"
)

func TestEmitItem_1(t *testing.T) {
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
	actual := l.emitItem(i)

	if actual != nil {
		t.Errorf("Expected nil, got %v", actual)
	}

	if l.item != expected {
		t.Errorf("Expected %v, got %v", expected, l.item)
	}
}
