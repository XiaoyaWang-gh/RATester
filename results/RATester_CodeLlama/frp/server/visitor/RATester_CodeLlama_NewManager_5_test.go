package visitor

import (
	"fmt"
	"testing"
)

func TestNewManager_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	vm := NewManager()
	if vm == nil {
		t.Fatal("NewManager() failed")
	}
}
