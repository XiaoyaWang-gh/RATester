package config

import (
	"fmt"
	"testing"
)

func TestDefaultBool_2(t *testing.T) {
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
	key := "key"
	defaultVal := false
	want := true
	got := c.DefaultBool(key, defaultVal)
	if got != want {
		t.Errorf("DefaultBool() = %v, want %v", got, want)
	}
}
