package codegen

import (
	"fmt"
	"testing"
)

func TestDeclaration_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Method{
		Name: "Foo",
	}
	if got := m.Declaration("m"); got != "func (m Method) Foo()" {
		t.Errorf("Declaration() = %v, want %v", got, "func (m Method) Foo()")
	}
}
