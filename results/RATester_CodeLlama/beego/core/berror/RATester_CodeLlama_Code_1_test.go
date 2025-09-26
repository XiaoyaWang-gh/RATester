package berror

import (
	"fmt"
	"testing"
)

func TestCode_1(t *testing.T) {
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
	if c.Code() != 1 {
		t.Errorf("Code() = %v, want %v", c.Code(), 1)
	}
}
