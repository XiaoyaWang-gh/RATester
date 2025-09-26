package codegen

import (
	"fmt"
	"testing"
)

func TestDeclarationNamed_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Method{
		Name: "Foo",
		In: []string{
			"string",
			"int",
			"interface{}",
			"net.Url",
		},
		Out: []string{
			"string",
			"int",
			"interface{}",
			"net.Url",
		},
	}
	receiver := "Foo"
	got := m.DeclarationNamed(receiver)
	want := "func (Foo Foo) Foo(string, int, interface{}, net.Url) (string, int, interface{}, net.Url)"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
