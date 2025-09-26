package web

import (
	"fmt"
	"testing"
)

func TestNewFlash_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fd := NewFlash()
	if fd == nil {
		t.Error("NewFlash() failed")
	}
}
