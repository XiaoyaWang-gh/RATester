package berror

import (
	"fmt"
	"testing"
)

func TestName_1(t *testing.T) {
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
	if c.Name() != "name" {
		t.Errorf("Name() = %v, want %v", c.Name(), "name")
	}
}
