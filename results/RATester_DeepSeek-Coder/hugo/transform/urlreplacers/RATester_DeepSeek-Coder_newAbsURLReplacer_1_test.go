package urlreplacers

import (
	"fmt"
	"testing"
)

func TestNewAbsURLReplacer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := newAbsURLReplacer()
	if r == nil {
		t.Errorf("newAbsURLReplacer() returned nil")
	}
	if len(r.htmlQuotes) != 2 {
		t.Errorf("newAbsURLReplacer() returned invalid htmlQuotes length, expected 2, got %d", len(r.htmlQuotes))
	}
	if len(r.xmlQuotes) != 2 {
		t.Errorf("newAbsURLReplacer() returned invalid xmlQuotes length, expected 2, got %d", len(r.xmlQuotes))
	}
}
