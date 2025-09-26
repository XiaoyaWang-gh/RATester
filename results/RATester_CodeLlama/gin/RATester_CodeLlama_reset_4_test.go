package gin

import (
	"fmt"
	"testing"
)

func TestReset_4(t *testing.T) {
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
	c.reset()
	if len(c.Params) != 0 {
		t.Errorf("Params should be empty")
	}
}
