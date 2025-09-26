package gin

import (
	"fmt"
	"testing"
)

func TestGetQueryArray_1(t *testing.T) {
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
	c.initQueryCache()
	values, ok := c.GetQueryArray("key")
	if !ok {
		t.Errorf("GetQueryArray() ok = %v, want %v", ok, true)
	}
	if len(values) != 1 {
		t.Errorf("GetQueryArray() len(values) = %v, want %v", len(values), 1)
	}
	if values[0] != "value" {
		t.Errorf("GetQueryArray() values[0] = %v, want %v", values[0], "value")
	}
}
