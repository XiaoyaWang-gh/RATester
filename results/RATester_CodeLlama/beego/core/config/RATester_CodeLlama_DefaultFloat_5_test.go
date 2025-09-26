package config

import (
	"fmt"
	"testing"
)

func TestDefaultFloat_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := "key"
	defaultVal := float64(1.0)
	want := float64(1.0)
	got := DefaultFloat(key, defaultVal)
	if got != want {
		t.Errorf("DefaultFloat() = %v, want %v", got, want)
	}
}
