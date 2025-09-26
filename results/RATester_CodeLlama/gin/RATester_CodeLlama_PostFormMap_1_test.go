package gin

import (
	"fmt"
	"testing"
)

func TestPostFormMap_1(t *testing.T) {
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
	dicts := c.PostFormMap("key")
	if dicts["key"] != "value" {
		t.Errorf("PostFormMap() = %v, want %v", dicts["key"], "value")
	}
}
