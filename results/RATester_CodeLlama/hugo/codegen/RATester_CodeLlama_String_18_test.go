package codegen

import (
	"fmt"
	"testing"
)

func TestString_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Method{
		Name: "Foo",
	}
	if m.String() != "Foo\n" {
		t.Errorf("expected %q, got %q", "Foo\n", m.String())
	}
}
