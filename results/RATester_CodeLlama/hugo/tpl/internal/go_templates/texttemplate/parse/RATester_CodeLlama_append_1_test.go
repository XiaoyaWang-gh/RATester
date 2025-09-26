package parse

import (
	"fmt"
	"testing"
)

func TestAppend_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &CommandNode{}
	arg := &IdentifierNode{}
	c.append(arg)
	if len(c.Args) != 1 {
		t.Errorf("c.Args = %v; want 1", len(c.Args))
	}
	if c.Args[0] != arg {
		t.Errorf("c.Args[0] = %v; want %v", c.Args[0], arg)
	}
}
