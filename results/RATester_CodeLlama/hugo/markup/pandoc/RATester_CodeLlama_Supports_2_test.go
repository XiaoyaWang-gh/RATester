package pandoc

import (
	"fmt"
	"testing"
)

func TestSupports_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if Supports() {
		t.Error("Expected false")
	}
}
