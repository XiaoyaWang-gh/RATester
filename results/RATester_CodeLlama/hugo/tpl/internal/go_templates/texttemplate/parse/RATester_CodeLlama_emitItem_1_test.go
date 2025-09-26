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
		input: `
			{{define "test"}}
				{{.}}
			{{end}}
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
			emitComment: false,
			breakOK:     false,
			continueOK:  false,
		},
	}
	l.emitItem(item{
		typ:  itemText,
		pos:  0,
		val:  "{{define \"test\"}}",
		line: 1,
	})
	if l.item.typ != itemText {
		t.Errorf("l.item.typ = %v, want %v", l.item.typ, itemText)
	}
	if l.item.pos != 0 {
		t.Errorf("l.item.pos = %v, want %v", l.item.pos, 0)
	}
	if l.item.val != "{{define \"test\"}}" {
		t.Errorf("l.item.val = %v, want %v", l.item.val, "{{define \"test\"}}")
	}
	if l.item.line != 1 {
		t.Errorf("l.item.line = %v, want %v", l.item.line, 1)
	}
}
