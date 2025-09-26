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

	height := d.Height()
	if height != 100 {
		t.Errorf("Expected height to be 100, got %d", height)
	}
}
