package config

import (
	"fmt"
	"strings"
	"testing"
)

func TestSet_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &fakeConfigContainer{
		data: make(map[string]string),
	}
	key := "key"
	val := "val"
	err := c.Set(key, val)
	if err != nil {
		t.Errorf("Set() error = %v", err)
		return
	}
	if c.data[strings.ToLower(key)] != val {
		t.Errorf("Set() data = %v, want %v", c.data[strings.ToLower(key)], val)
	}
}
