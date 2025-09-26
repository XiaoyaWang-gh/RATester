package config

import (
	"fmt"
	"testing"
)

func TestUnmarshaler_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &fakeConfigContainer{}
	prefix := "prefix"
	obj := "obj"
	opt := []DecodeOption{}
	err := c.Unmarshaler(prefix, obj, opt...)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
