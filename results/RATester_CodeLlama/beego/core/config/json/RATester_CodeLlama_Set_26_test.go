package json

import (
	"fmt"
	"testing"
)

func TestSet_26(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &JSONConfigContainer{
		data: make(map[string]interface{}),
	}
	key := "key"
	val := "val"
	err := c.Set(key, val)
	if err != nil {
		t.Errorf("Set failed. Error: %v", err)
	}
	if c.data[key] != val {
		t.Errorf("Set failed. Expected: %v, actual: %v", val, c.data[key])
	}
}
