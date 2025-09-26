package config

import (
	"fmt"
	"testing"
)

func TestNew_65(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := New()
	if p == nil {
		t.Error("New() returned nil")
	}
}
