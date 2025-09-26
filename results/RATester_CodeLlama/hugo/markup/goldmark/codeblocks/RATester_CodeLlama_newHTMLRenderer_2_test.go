package codeblocks

import (
	"fmt"
	"testing"
)

func TestNewHTMLRenderer_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := newHTMLRenderer()
	if r == nil {
		t.Error("newHTMLRenderer() returned nil")
	}
}
