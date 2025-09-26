package helpers

import (
	"fmt"
	"testing"
)

func TestString_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := NamedSlice{
		Name:  "foo",
		Slice: []string{"bar", "baz"},
	}
	if n.String() != "foo/bar,baz" {
		t.Errorf("expected %q, got %q", "foo/bar,baz", n.String())
	}
}
