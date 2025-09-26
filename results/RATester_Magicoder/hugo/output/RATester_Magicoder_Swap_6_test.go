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
		{Name: "Format1"},
		{Name: "Format2"},
	}

	formats.Swap(0, 1)

	if formats[0].Name != "Format2" || formats[1].Name != "Format1" {
		t.Errorf("Swap failed. Expected: Format2, Format1. Got: %s, %s", formats[0].Name, formats[1].Name)
	}
}
