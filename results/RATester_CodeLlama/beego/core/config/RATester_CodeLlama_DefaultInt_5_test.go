package config

import (
	"fmt"
	"testing"
)

func TestDefaultInt_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &BaseConfiger{}
	key := "key"
	defaultVal := 1
	res := c.DefaultInt(key, defaultVal)
	if res != defaultVal {
		t.Errorf("DefaultInt() = %v, want %v", res, defaultVal)
	}
}
