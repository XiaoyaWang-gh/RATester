package diagrams

import (
	"fmt"
	"testing"

	"github.com/bep/goat"
)

func TestWidth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := goatDiagram{
		d: goat.SVG{
			Width: 100,
		},
	}

	expected := 100
	actual := d.Width()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
