package berror

import (
	"fmt"
	"testing"
)

func TestDesc_1(t *testing.T) {
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
	if c.Desc() != "desc" {
		t.Errorf("Desc() = %v, want %v", c.Desc(), "desc")
	}
}
