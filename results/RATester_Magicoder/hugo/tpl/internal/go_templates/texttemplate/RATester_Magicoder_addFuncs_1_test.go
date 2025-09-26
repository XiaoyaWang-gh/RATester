package template

import (
	"fmt"
	"testing"
)

func TestaddFuncs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	out := make(FuncMap)
	in := make(FuncMap)
	in["testFunc"] = func() {}

	addFuncs(out, in)

	if _, ok := out["testFunc"]; !ok {
		t.Error("Expected testFunc to be in out")
	}
}
