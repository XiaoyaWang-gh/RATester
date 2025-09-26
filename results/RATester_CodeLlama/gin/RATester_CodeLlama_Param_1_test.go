package gin

import (
	"fmt"
	"testing"
)

func TestParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Params: Params{
			{
				Key:   "key",
				Value: "value",
			},
		},
	}

	if c.Param("key") != "value" {
		t.Errorf("Expected value to be %s, but got %s", "value", c.Param("key"))
	}
}
