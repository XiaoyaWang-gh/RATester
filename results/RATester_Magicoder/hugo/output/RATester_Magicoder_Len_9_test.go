package output

import (
	"fmt"
	"testing"
)

func TestLen_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	formats := Formats{
		{Name: "Format1"},
		{Name: "Format2"},
		{Name: "Format3"},
	}

	expected := 3
	actual := formats.Len()

	if actual != expected {
		t.Errorf("Expected Len() to return %d, but got %d", expected, actual)
	}
}
