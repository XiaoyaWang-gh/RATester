package tcp

import (
	"fmt"
	"testing"
)

func TestLess_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	routes := routes{
		{priority: 1},
		{priority: 2},
		{priority: 3},
	}

	if routes.Less(0, 1) {
		t.Error("Expected false, got true")
	}

	if !routes.Less(1, 0) {
		t.Error("Expected true, got false")
	}
}
