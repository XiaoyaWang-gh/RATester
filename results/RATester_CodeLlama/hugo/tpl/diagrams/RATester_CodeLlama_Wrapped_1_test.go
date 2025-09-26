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
		d: goat.SVG{
			Width:  100,
			Height: 100,
		},
	}
	want := template.HTML(`<svg width="100" height="100"></svg>`)
	if got := d.Wrapped(); got != want {
		t.Errorf("Wrapped() = %v, want %v", got, want)
	}
}
