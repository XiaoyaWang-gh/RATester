package parse

import (
	"fmt"
	"testing"
)

func TestLex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := lex("", "", "", "")
	if l.name != "" {
		t.Errorf("l.name = %q, want \"\"", l.name)
	}
	if l.input != "" {
		t.Errorf("l.input = %q, want \"\"", l.input)
	}
	if l.leftDelim != leftDelim {
		t.Errorf("l.leftDelim = %q, want %q", l.leftDelim, leftDelim)
	}
	if l.rightDelim != rightDelim {
		t.Errorf("l.rightDelim = %q, want %q", l.rightDelim, rightDelim)
	}
	if l.line != 1 {
		t.Errorf("l.line = %d, want 1", l.line)
	}
	if l.startLine != 1 {
		t.Errorf("l.startLine = %d, want 1", l.startLine)
	}
	if l.insideAction {
		t.Errorf("l.insideAction = true, want false")
	}
}
