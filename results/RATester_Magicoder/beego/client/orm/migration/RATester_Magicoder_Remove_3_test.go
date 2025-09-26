package migration

import (
	"fmt"
	"testing"
)

func TestRemove_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	column := &Column{
		Name: "test_column",
	}

	column.Remove()

	if !column.remove {
		t.Error("Expected column.remove to be true, but it was false")
	}
}
