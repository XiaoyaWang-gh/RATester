package gin

import (
	"fmt"
	"testing"
)

func TestPostFormArray_1(t *testing.T) {
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
	values := c.PostFormArray("key")
	if len(values) != 1 {
		t.Errorf("PostFormArray() = %v, want %v", len(values), 1)
	}
	if values[0] != "value" {
		t.Errorf("PostFormArray() = %v, want %v", values[0], "value")
	}
}
