package maps

import (
	"fmt"
	"testing"
)

func TestGet_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Scratch{
		values: map[string]any{
			"key1": "value1",
			"key2": "value2",
		},
	}

	key := "key1"

	val := c.Get(key)

	if val != "value1" {
		t.Errorf("Expected value1, got %v", val)
	}
}
