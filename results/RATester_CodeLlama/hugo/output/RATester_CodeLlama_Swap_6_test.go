package output

import (
	"fmt"
	"testing"
)

func TestSwap_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	formats := Formats{
		{Name: "a"},
		{Name: "b"},
		{Name: "c"},
	}

	formats.Swap(0, 1)

	if formats[0].Name != "b" {
		t.Errorf("Expected formats[0].Name to be 'b', got %q", formats[0].Name)
	}

	if formats[1].Name != "a" {
		t.Errorf("Expected formats[1].Name to be 'a', got %q", formats[1].Name)
	}
}
