package config

import (
	"fmt"
	"testing"
)

func TestInt64_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &fakeConfigContainer{
		data: map[string]string{
			"key": "123",
		},
	}
	v, err := c.Int64("key")
	if err != nil {
		t.Errorf("Int64() error = %v", err)
		return
	}
	if v != 123 {
		t.Errorf("Int64() = %v, want %v", v, 123)
	}
}
