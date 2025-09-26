package parse

import (
	"fmt"
	"testing"
)

func Testaccept_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &lexer{
		name:  "test",
		input: "test",
	}

	valid := "test"
	expected := true
	result := l.accept(valid)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
