package page

import (
	"fmt"
	"html/template"
	"testing"
)

func TestPrintableValue_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := Summary{
		Text:      template.HTML("foo"),
		Type:      "auto",
		Truncated: true,
	}

	if got := s.PrintableValue(); got != "foo" {
		t.Errorf("Summary.PrintableValue() = %v, want %v", got, "foo")
	}
}
