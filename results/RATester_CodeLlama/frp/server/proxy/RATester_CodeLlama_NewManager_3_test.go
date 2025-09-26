package proxy

import (
	"fmt"
	"testing"
)

func TestNewManager_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pm := NewManager()
	if pm == nil {
		t.Error("NewManager() failed")
	}
}
