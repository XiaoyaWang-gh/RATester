package config

import (
	"fmt"
	"testing"
)

func TestDefaultFloat_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &BaseConfiger{}
	key := "key"
	defaultVal := float64(1)
	res := c.DefaultFloat(key, defaultVal)
	if res != defaultVal {
		t.Errorf("DefaultFloat() = %v, want %v", res, defaultVal)
	}
}
