package diagrams

import (
	"fmt"
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
			Body: "test",
		},
	}
	if got := d.Inner(); got != "test" {
		t.Errorf("goatDiagram.Inner() = %v, want %v", got, "test")
	}
}
