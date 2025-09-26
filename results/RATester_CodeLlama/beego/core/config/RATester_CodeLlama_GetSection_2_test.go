package config

import (
	"fmt"
	"testing"
)

func TestGetSection_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &fakeConfigContainer{
		data: map[string]string{
			"section1": "value1",
			"section2": "value2",
		},
	}
	section, err := c.GetSection("section1")
	if err != nil {
		t.Errorf("GetSection() error = %v", err)
		return
	}
	if len(section) != 1 {
		t.Errorf("GetSection() = %v, want %v", len(section), 1)
		return
	}
	if section["section1"] != "value1" {
		t.Errorf("GetSection() = %v, want %v", section["section1"], "value1")
		return
	}
}
