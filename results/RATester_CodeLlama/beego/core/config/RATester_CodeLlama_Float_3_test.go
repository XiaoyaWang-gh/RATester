package config

import (
	"fmt"
	"testing"
)

func TestFloat_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &fakeConfigContainer{
		data: map[string]string{
			"key": "1.23",
		},
	}
	v, err := c.Float("key")
	if err != nil {
		t.Errorf("Float() error = %v", err)
		return
	}
	if v != 1.23 {
		t.Errorf("Float() = %v, want %v", v, 1.23)
	}
}
