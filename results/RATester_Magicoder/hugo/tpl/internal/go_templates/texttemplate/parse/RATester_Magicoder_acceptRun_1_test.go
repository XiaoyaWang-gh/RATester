package parse

import (
	"fmt"
	"testing"
)

func TestacceptRun_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &lexer{
		input: "abcdefg",
	}
	l.acceptRun("abc")
	if l.input != "defg" {
		t.Errorf("Expected input to be 'defg', but got '%s'", l.input)
	}
}
