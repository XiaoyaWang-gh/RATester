package diagrams

import (
	"fmt"
	"testing"
)

func TestWidth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := goatDiagram{}
	if d.Width() != 0 {
		t.Errorf("Expected 0, got %d", d.Width())
	}
}
