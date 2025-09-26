package gin

import (
	"fmt"
	"testing"
)

func TestGetBool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Keys: map[string]any{
			"key": true,
		},
	}
	b := c.GetBool("key")
	if b != true {
		t.Errorf("GetBool() = %v, want %v", b, true)
	}
}
