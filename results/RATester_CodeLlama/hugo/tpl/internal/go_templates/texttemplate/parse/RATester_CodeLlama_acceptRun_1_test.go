package parse

import (
	"fmt"
	"testing"
)

func TestAcceptRun_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &lexer{
		name: "test",
		input: `
			{{range .}}
				{{.Name}}
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
	l.acceptRun("abc")
	if l.pos != 3 {
		t.Errorf("l.pos = %d, want 3", l.pos)
	}
	if l.input[l.pos] != 'a' {
		t.Errorf("l.input[l.pos] = %c, want 'a'", l.input[l.pos])
	}
}
