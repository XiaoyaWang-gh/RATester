package plugin

import (
	"fmt"
	"testing"
)

func TestNewManager_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := NewManager()
	if m == nil {
		t.Fatal("NewManager() failed")
	}
}
