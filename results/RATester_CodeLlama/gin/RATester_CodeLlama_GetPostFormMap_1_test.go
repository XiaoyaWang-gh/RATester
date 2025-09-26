package gin

import (
	"fmt"
	"testing"
)

func TestGetPostFormMap_1(t *testing.T) {
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
	m, ok := c.GetPostFormMap("key")
	if !ok {
		t.Errorf("GetPostFormMap() = %v, want %v", m, "value")
	}
}
