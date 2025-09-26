package loggers

import (
	"fmt"
	"strings"
	"testing"
)

func TestErrors_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{
		errors: &strings.Builder{},
	}

	// Test case 1: No errors
	l.errors.Reset()
	if got := l.Errors(); got != "" {
		t.Errorf("Errors() = %v, want \"\"", got)
	}

	// Test case 2: With errors
	l.errors.WriteString("Error 1\nError 2\n")
	want := "Error 1\nError 2\n"
	if got := l.Errors(); got != want {
		t.Errorf("Errors() = %v, want %v", got, want)
	}
}
