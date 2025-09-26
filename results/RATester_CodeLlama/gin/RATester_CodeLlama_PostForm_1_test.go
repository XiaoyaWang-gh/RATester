package gin

import (
	"fmt"
	"testing"
)

func TestPostForm_1(t *testing.T) {
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
	if c.PostForm("key") != "value" {
		t.Errorf("PostForm() = %v, want %v", c.PostForm("key"), "value")
	}
}
