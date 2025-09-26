package es

import (
	"fmt"
	"testing"
)

func TestNewES_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cw := NewES()
	if cw == nil {
		t.Errorf("NewES() = %v, want %v", cw, nil)
	}
}
