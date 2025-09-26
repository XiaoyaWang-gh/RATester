package berror

import (
	"fmt"
	"testing"
)

func TestGet_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cr := &codeRegistry{
		codes: map[uint32]*codeDefinition{
			1: {
				code:   1,
				module: "module1",
				desc:   "desc1",
				name:   "name1",
			},
			2: {
				code:   2,
				module: "module2",
				desc:   "desc2",
				name:   "name2",
			},
		},
	}
	code := uint32(1)
	c, ok := cr.Get(code)
	if !ok {
		t.Errorf("code %d not found", code)
	}
	if c.Code() != code {
		t.Errorf("code %d not equal to %d", c.Code(), code)
	}
	if c.Module() != "module1" {
		t.Errorf("module %s not equal to %s", c.Module(), "module1")
	}
	if c.Desc() != "desc1" {
		t.Errorf("desc %s not equal to %s", c.Desc(), "desc1")
	}
	if c.Name() != "name1" {
		t.Errorf("name %s not equal to %s", c.Name(), "name1")
	}
}
