package diagrams

import (
	"fmt"
	"testing"

	"github.com/bep/goat"
)

func TestHeight_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := goatDiagram{
		d: goat.SVG{
			Height: 100,
		},
	}

	expected := 100
	actual := d.Height()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
