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
		input: "abc123",
	}

	l.acceptRun("abc")

	if l.input != "123" {
		t.Errorf("Expected input to be '123', got '%s'", l.input)
	}
}
