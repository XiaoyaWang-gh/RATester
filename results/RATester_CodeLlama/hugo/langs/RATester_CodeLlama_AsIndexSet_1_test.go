package langs

import (
	"fmt"
	"testing"
)

func TestAsIndexSet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := Languages{
		{Lang: "en"},
		{Lang: "no"},
		{Lang: "en"},
	}

	m := l.AsIndexSet()

	if len(m) != 2 {
		t.Errorf("Expected 2 languages, got %d", len(m))
	}

	if m["en"] != 0 {
		t.Errorf("Expected en to be at index 0, got %d", m["en"])
	}

	if m["no"] != 1 {
		t.Errorf("Expected no to be at index 1, got %d", m["no"])
	}
}
