package config

import (
	"fmt"
	"testing"
)

func TestSet_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &IniConfigContainer{}
	c.Set("key", "val")
}
