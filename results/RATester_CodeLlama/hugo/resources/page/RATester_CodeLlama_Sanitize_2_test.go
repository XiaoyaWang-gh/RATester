package page

import (
	"fmt"
	"testing"
)

func TestSanitize_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pagePathBuilder{
		els: []string{"a", "b", "c"},
	}
	p.Sanitize()
	if p.els[0] != "a" {
		t.Errorf("Expected %q, got %q", "a", p.els[0])
	}
	if p.els[1] != "b" {
		t.Errorf("Expected %q, got %q", "b", p.els[1])
	}
	if p.els[2] != "c" {
		t.Errorf("Expected %q, got %q", "c", p.els[2])
	}
}
