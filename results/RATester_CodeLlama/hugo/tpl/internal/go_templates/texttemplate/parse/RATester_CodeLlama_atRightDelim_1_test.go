package parse

import (
	"fmt"
	"testing"
)

func TestAtRightDelim_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &lexer{
		input: "{{ }}",
	}
	delim, trimSpaces := l.atRightDelim()
	if delim != true {
		t.Errorf("delim = %v, want %v", delim, true)
	}
	if trimSpaces != false {
		t.Errorf("trimSpaces = %v, want %v", trimSpaces, false)
	}
}
