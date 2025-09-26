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
			Height: 10,
		},
	}
	if d.Height() != 10 {
		t.Errorf("Height() = %d, want %d", d.Height(), 10)
	}
}
