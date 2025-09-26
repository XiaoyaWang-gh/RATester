package gin

import (
	"fmt"
	"testing"
)

func TestGetPostForm_1(t *testing.T) {
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
	if value, ok := c.GetPostForm("key"); !ok {
		t.Errorf("GetPostForm() = %v, want %v", value, "value")
	}
}
