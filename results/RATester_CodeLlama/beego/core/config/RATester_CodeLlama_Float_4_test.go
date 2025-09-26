package config

import (
	"fmt"
	"testing"
)

func TestFloat_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &IniConfigContainer{}
	c.data = map[string]map[string]string{
		"section": {
			"key": "1.23",
		},
	}
	v, err := c.Float("section:key")
	if err != nil {
		t.Error(err)
	}
	if v != 1.23 {
		t.Errorf("want 1.23, but %f", v)
	}
}
