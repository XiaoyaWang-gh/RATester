package gin

import (
	"fmt"
	"testing"
)

func TestGetStringMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Keys: map[string]any{
			"key": map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
		},
	}
	sm := c.GetStringMap("key")
	if sm == nil {
		t.Errorf("GetStringMap() = %v, want %v", sm, c.Keys["key"])
	}
}
