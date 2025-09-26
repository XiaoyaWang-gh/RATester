package berror

import (
	"fmt"
	"testing"
)

func TestModule_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &codeDefinition{
		code:   1,
		module: "module",
		desc:   "desc",
		name:   "name",
	}
	if c.Module() != "module" {
		t.Errorf("Module() = %v, want %v", c.Module(), "module")
	}
}
