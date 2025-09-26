package scss

import (
	"fmt"
	"testing"
)

func TestSupports_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	result := Supports()
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}
