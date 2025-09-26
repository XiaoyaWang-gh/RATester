package page

import (
	"fmt"
	"testing"
)

func TestPrintableValue_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	summary := Summary{
		Text:      "Test Summary",
		Type:      "auto",
		Truncated: false,
	}

	expected := "Test Summary"
	result := summary.PrintableValue()

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
