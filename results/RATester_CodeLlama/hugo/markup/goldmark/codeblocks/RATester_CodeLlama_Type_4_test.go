package codeblocks

import (
	"fmt"
	"testing"
)

func TestType_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &codeBlockContext{
		lang: "go",
	}
	if c.Type() != "go" {
		t.Errorf("Type() = %v, want %v", c.Type(), "go")
	}
}
