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
		{Name: "format1"},
		{Name: "format2"},
		{Name: "format3"},
	}

	expected := 3
	actual := formats.Len()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
