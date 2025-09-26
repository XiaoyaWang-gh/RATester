package parse

import (
	"fmt"
	"testing"
)

func TestNew_27(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "test"
	funcs := []map[string]any{
		{
			"func1": func() {
				fmt.Println("func1")
			},
		},
		{
			"func2": func() {
				fmt.Println("func2")
			},
		},
	}
	tree := New(name, funcs...)
	if tree.Name != name {
		t.Errorf("tree.Name = %v, want %v", tree.Name, name)
	}
	if len(tree.funcs) != len(funcs) {
		t.Errorf("len(tree.funcs) = %v, want %v", len(tree.funcs), len(funcs))
	}
	for i, f := range tree.funcs {
		if len(f) != len(funcs[i]) {
			t.Errorf("len(f) = %v, want %v", len(f), len(funcs[i]))
		}
		for k, v := range f {
			if funcs[i][k] != v {
				t.Errorf("f[%v] = %v, want %v", k, v, funcs[i][k])
			}
		}
	}
}
