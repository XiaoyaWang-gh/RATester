package config

import (
	"fmt"
	"testing"
)

func TestDIY_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &fakeConfigContainer{
		data: map[string]string{
			"key": "value",
		},
	}
	v, err := c.DIY("key")
	if err != nil {
		t.Errorf("DIY() error = %v", err)
		return
	}
	if v != "value" {
		t.Errorf("DIY() = %v, want %v", v, "value")
	}
}
