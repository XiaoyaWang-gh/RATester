package gin

import (
	"fmt"
	"testing"
)

func TestGetPostFormArray_1(t *testing.T) {
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
	values, ok := c.GetPostFormArray("key")
	if !ok {
		t.Errorf("GetPostFormArray() ok = %v, want %v", ok, true)
	}
	if len(values) != 1 {
		t.Errorf("GetPostFormArray() len(values) = %v, want %v", len(values), 1)
	}
	if values[0] != "value" {
		t.Errorf("GetPostFormArray() values[0] = %v, want %v", values[0], "value")
	}
}
