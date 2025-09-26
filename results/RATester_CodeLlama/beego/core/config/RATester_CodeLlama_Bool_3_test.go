package config

import (
	"fmt"
	"testing"
)

func TestBool_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &fakeConfigContainer{
		data: map[string]string{
			"key": "true",
		},
	}
	b, err := c.Bool("key")
	if err != nil {
		t.Errorf("Bool() error = %v", err)
		return
	}
	if !b {
		t.Errorf("Bool() = %v, want %v", b, true)
	}
}
