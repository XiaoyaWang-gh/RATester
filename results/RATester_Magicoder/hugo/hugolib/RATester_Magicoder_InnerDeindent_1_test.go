package hugolib

import (
	"fmt"
	"html/template"
	"sync"
	"testing"
)

func TestInnerDeindent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scp := &ShortcodeWithPage{
		Inner:             template.HTML("  Hello, World!"),
		indentation:       "  ",
		innerDeindentInit: sync.Once{},
	}

	expected := template.HTML("Hello, World!")
	result := scp.InnerDeindent()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
