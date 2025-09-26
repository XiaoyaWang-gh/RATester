package config

import (
	"fmt"
	"testing"
)

func TestDefaultFloat_2(t *testing.T) {
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
	key := "key"
	defaultVal := 1.234
	want := 1.23
	got := c.DefaultFloat(key, defaultVal)
	if got != want {
		t.Errorf("DefaultFloat() = %v, want %v", got, want)
	}
}
