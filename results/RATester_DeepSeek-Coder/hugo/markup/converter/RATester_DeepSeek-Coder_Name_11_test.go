package converter

import (
	"fmt"
	"testing"
)

func TestName_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := newConverter{
		name: "test",
	}

	result := n.Name()

	if result != "test" {
		t.Errorf("Expected 'test', got '%s'", result)
	}
}
