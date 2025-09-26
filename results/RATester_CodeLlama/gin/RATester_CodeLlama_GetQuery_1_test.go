package gin

import (
	"fmt"
	"testing"
)

func TestGetQuery_1(t *testing.T) {
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

	if v, ok := c.GetQuery("key"); !ok {
		t.Errorf("Expected value to be %v, got %v", "value", v)
	}

	if v, ok := c.GetQuery("key2"); ok {
		t.Errorf("Expected value to be %v, got %v", "", v)
	}
}
