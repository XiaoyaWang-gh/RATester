package maps

import (
	"fmt"
	"testing"
)

func TestSet_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Scratch{}
	key := "key"
	value := "value"
	c.Set(key, value)
	if c.values[key] != value {
		t.Errorf("Set failed")
	}
}
