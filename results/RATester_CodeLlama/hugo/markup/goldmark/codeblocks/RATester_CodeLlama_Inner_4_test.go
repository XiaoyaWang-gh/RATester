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

	c := &codeBlockContext{}
	c.code = "code"
	if c.Inner() != "code" {
		t.Errorf("codeBlockContext.Inner() = %v, want %v", c.Inner(), "code")
	}
}
