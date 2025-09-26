package config

import (
	"fmt"
	"testing"
)

func TestIsZero_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Redirect{}
	if !r.IsZero() {
		t.Errorf("IsZero() = %v, want %v", r.IsZero(), true)
	}
}
