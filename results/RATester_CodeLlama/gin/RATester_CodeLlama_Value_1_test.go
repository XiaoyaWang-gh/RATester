package gin

import (
	"fmt"
	"testing"
)

func TestValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Keys: map[string]any{
			"key1": "value1",
			"key2": "value2",
		},
	}

	if c.Value("key1") != "value1" {
		t.Errorf("Value() = %v, want %v", c.Value("key1"), "value1")
	}

	if c.Value("key2") != "value2" {
		t.Errorf("Value() = %v, want %v", c.Value("key2"), "value2")
	}

	if c.Value("key3") != nil {
		t.Errorf("Value() = %v, want %v", c.Value("key3"), nil)
	}
}
