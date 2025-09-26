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

	n := newConverter{name: "test"}
	if n.Name() != "test" {
		t.Errorf("Expected name to be 'test', got '%s'", n.Name())
	}
}
