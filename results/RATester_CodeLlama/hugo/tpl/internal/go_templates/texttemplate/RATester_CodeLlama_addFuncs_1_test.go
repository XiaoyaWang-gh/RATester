package template

import (
	"fmt"
	"testing"
)

func TestAddFuncs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	out := make(FuncMap)
	in := make(FuncMap)
	addFuncs(out, in)
	if len(out) != 0 {
		t.Errorf("out should be empty")
	}
	in["test"] = func() {}
	addFuncs(out, in)
	if len(out) != 1 {
		t.Errorf("out should have 1 element")
	}
	if out["test"] != in["test"] {
		t.Errorf("out should have the same function as in")
	}
}
