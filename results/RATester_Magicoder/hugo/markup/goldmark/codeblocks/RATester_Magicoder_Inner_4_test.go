package codeblocks

import (
	"fmt"
	"testing"
)

func TestInner_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &codeBlockContext{
		lang: "go",
		code: "func TestInner(t *testing.T) {...}",
	}

	expected := "func TestInner(t *testing.T) {...}"
	actual := ctx.Inner()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
