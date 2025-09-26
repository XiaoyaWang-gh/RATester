package config

import (
	"fmt"
	"testing"
)

func TestDefaultString_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &BaseConfiger{}
	key := "key"
	defaultVal := "defaultVal"
	res := c.DefaultString(key, defaultVal)
	if res != defaultVal {
		t.Errorf("DefaultString() = %v, want %v", res, defaultVal)
	}
}
