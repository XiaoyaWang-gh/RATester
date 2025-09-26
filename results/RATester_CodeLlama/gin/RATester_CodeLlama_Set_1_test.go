package gin

import (
	"fmt"
	"testing"
)

func TestSet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{}
	c.Set("key", "value")
	if c.Keys["key"] != "value" {
		t.Errorf("Expected value to be 'value', got '%v'", c.Keys["key"])
	}
}
