package gin

import (
	"fmt"
	"testing"
)

func TestGetString_1(t *testing.T) {
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
	if s := c.GetString("key"); s != "value" {
		t.Errorf("GetString() = %v, want %v", s, "value")
	}
}
