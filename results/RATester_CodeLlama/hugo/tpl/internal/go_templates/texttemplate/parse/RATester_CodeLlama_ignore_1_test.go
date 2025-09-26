package parse

import (
	"fmt"
	"testing"
)

func TestIgnore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &lexer{
		name: "test",
		input: `
			{{- "hello" -}}
			{{- "world" -}}
		`,
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
		options: lexOptions{
			emitComment: true,
			breakOK:     true,
			continueOK:  true,
		},
	}
	l.ignore()
	if l.line != 2 {
		t.Errorf("l.line = %d, want 2", l.line)
	}
	if l.start != 12 {
		t.Errorf("l.start = %d, want 12", l.start)
	}
	if l.startLine != 2 {
		t.Errorf("l.startLine = %d, want 2", l.startLine)
	}
}
