package config

import (
	"fmt"
	"testing"
)

func TestStrings_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &IniConfigContainer{}
	c.data = make(map[string]map[string]string)
	c.data["section"] = make(map[string]string)
	c.data["section"]["key"] = "val"
	c.keyComment = make(map[string]string)
	c.keyComment["section:key"] = "comment"
	v, err := c.Strings("section:key")
	if err != nil {
		t.Error(err)
	}
	if v[0] != "val" {
		t.Error("val is not equal")
	}
	if v[1] != "comment" {
		t.Error("comment is not equal")
	}
}
