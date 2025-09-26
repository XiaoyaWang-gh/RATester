package diagrams

import (
	"fmt"
	"html/template"
	"testing"

	"github.com/bep/goat"
)

func TestWrapped_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := goatDiagram{
		d: goat.SVG{},
	}
	expected := template.HTML(d.d.String())
	result := d.Wrapped()
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
