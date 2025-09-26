package config

import (
	"fmt"
	"testing"
)

func TestDefaultInt64_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &IniConfigContainer{}
	c.data = make(map[string]map[string]string)
	c.data["section"] = make(map[string]string)
	c.data["section"]["key"] = "1"
	if v := c.DefaultInt64("section:key", 10); v != 1 {
		t.Errorf("DefaultInt64() = %v, want %v", v, 1)
	}
	if v := c.DefaultInt64("section:key1", 10); v != 10 {
		t.Errorf("DefaultInt64() = %v, want %v", v, 10)
	}
}
