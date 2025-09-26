package codegen

import (
	"fmt"
	"testing"
)

func TestDelegate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Method{
		Name: "Foo",
		In:   []string{"string", "int"},
		Out:  []string{"string"},
	}
	if m.Delegate("", "") != "return Foo(string, int) string" {
		t.Errorf("Delegate failed: %s", m.Delegate("", ""))
	}
}
