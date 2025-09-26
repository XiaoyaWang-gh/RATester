package tplimpl

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestnewTemplateExecuter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &deps.Deps{
		// Initialize your dependencies here
	}

	exe, funcs := newTemplateExecuter(d)

	// Test the function
	if exe == nil {
		t.Error("Expected a non-nil Executer, but got nil")
	}

	// Test the function
	if funcs == nil {
		t.Error("Expected a non-nil map of functions, but got nil")
	}
}
