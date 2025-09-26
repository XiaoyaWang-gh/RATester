package web

import (
	"fmt"
	"testing"
)

func TestNewControllerRegister_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := NewControllerRegister()
	if ctrl == nil {
		t.Error("NewControllerRegister() should not return nil")
	}
}
