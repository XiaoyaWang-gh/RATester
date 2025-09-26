package diagrams

import (
	"fmt"
	"html/template"
	"testing"

	"github.com/bep/goat"
)

func TestInner_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := goatDiagram{
		d: goat.SVG{
			Body: "<svg>...</svg>",
		},
	}
	expected := template.HTML(d.d.Body)
	result := d.Inner()
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
