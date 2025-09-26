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
		{Name: "format1"},
		{Name: "format2"},
	}

	formats.Swap(0, 1)

	if formats[0].Name != "format2" || formats[1].Name != "format1" {
		t.Errorf("Expected formats to be swapped, but they were not.")
	}
}
