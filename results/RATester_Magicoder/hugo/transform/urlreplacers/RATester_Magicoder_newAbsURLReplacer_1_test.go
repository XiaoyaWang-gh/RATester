package urlreplacers

import (
	"fmt"
	"testing"
)

func TestnewAbsURLReplacer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := newAbsURLReplacer()
	if r == nil {
		t.Error("Expected a non-nil absURLReplacer, got nil")
	}
	if len(r.htmlQuotes) != 2 {
		t.Errorf("Expected 2 htmlQuotes, got %d", len(r.htmlQuotes))
	}
	if len(r.xmlQuotes) != 2 {
		t.Errorf("Expected 2 xmlQuotes, got %d", len(r.xmlQuotes))
	}
}
