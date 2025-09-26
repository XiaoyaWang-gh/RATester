package hugolib

import (
	"fmt"
	"testing"
)

func TestGet_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &contentTreeReverseIndex{}
	key := "key"
	c.Get(key)
	if c.m[key] != nil {
		t.Errorf("c.m[key] = %v, want nil", c.m[key])
	}
}
