package gin

import (
	"fmt"
	"testing"
)

func TestMustGet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Keys: map[string]any{
			"key": "value",
		},
	}
	if c.MustGet("key") != "value" {
		t.Errorf("MustGet() = %v, want %v", c.MustGet("key"), "value")
	}
}
