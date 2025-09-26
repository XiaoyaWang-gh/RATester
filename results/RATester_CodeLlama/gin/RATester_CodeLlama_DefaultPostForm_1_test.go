package gin

import (
	"fmt"
	"testing"
)

func TestDefaultPostForm_1(t *testing.T) {
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
	if c.DefaultPostForm("key", "default") != "value" {
		t.Errorf("DefaultPostForm() = %v, want %v", c.DefaultPostForm("key", "default"), "value")
	}
	if c.DefaultPostForm("key2", "default") != "default" {
		t.Errorf("DefaultPostForm() = %v, want %v", c.DefaultPostForm("key2", "default"), "default")
	}
}
