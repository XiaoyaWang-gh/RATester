package parse

import (
	"fmt"
	"testing"
)

func Testignore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &lexer{
		name:         "test",
		input:        "test input",
		leftDelim:    "{{",
		rightDelim:   "}}",
		pos:          0,
		start:        0,
		atEOF:        false,
		parenDepth:   0,
		line:         1,
		startLine:    1,
		item:         item{},
		insideAction: false,
		options:      lexOptions{},
	}

	l.ignore()

	if l.line != 1 {
		t.Errorf("Expected line to be 1, but got %d", l.line)
	}
	if l.start != 4 {
		t.Errorf("Expected start to be 4, but got %d", l.start)
	}
	if l.startLine != 1 {
		t.Errorf("Expected startLine to be 1, but got %d", l.startLine)
	}
}
