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

	flash := NewFlash()
	if flash == nil {
		t.Error("Expected a FlashData, got nil")
	}
	if flash.Data == nil {
		t.Error("Expected a map, got nil")
	}
}
