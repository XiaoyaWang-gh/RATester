package media

import (
	"fmt"
	"testing"
)

func TestLess_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	types := Types{
		{Type: "application/json"},
		{Type: "text/html"},
		{Type: "image/png"},
	}

	if !types.Less(0, 1) {
		t.Error("Expected types[0] to be less than types[1]")
	}

	if types.Less(1, 0) {
		t.Error("Expected types[1] not to be less than types[0]")
	}

	if types.Less(0, 2) {
		t.Error("Expected types[0] not to be less than types[2]")
	}

	if !types.Less(2, 0) {
		t.Error("Expected types[2] to be less than types[0]")
	}
}
