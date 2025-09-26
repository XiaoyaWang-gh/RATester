package gin

import (
	"fmt"
	"testing"
)

func TestQuery_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Params: Params{
			Param{
				Key:   "key",
				Value: "value",
			},
		},
	}
	if c.Query("key") != "value" {
		t.Errorf("Query() = %v, want %v", c.Query("key"), "value")
	}
}
