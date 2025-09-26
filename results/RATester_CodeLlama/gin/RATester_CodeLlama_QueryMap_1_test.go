package gin

import (
	"fmt"
	"testing"
)

func TestQueryMap_1(t *testing.T) {
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
	dicts := c.QueryMap("key")
	if dicts["key"] != "value" {
		t.Errorf("QueryMap() = %v, want %v", dicts["key"], "value")
	}
}
