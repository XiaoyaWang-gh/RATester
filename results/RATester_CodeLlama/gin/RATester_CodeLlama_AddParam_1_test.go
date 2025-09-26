package gin

import (
	"fmt"
	"testing"
)

func TestAddParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Params: Params{},
	}
	c.AddParam("key", "value")
	if len(c.Params) != 1 {
		t.Errorf("Params length is not 1")
	}
	if c.Params[0].Key != "key" {
		t.Errorf("Params[0].Key is not key")
	}
	if c.Params[0].Value != "value" {
		t.Errorf("Params[0].Value is not value")
	}
}
