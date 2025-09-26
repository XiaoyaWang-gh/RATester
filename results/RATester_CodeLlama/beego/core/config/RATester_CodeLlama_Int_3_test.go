package config

import (
	"fmt"
	"testing"
)

func TestInt_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &IniConfigContainer{}
	c.Set("section:key", "1")
	v, err := c.Int("section:key")
	if err != nil {
		t.Error(err)
	}
	if v != 1 {
		t.Errorf("expect 1, but %d", v)
	}
}
