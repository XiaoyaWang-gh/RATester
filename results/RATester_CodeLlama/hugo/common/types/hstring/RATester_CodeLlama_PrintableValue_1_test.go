package hstring

import (
	"fmt"
	"html/template"
	"testing"
)

func TestPrintableValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := HTML("foo")
	if got := s.PrintableValue(); got != template.HTML("foo") {
		t.Errorf("PrintableValue() = %v, want %v", got, template.HTML("foo"))
	}
}
