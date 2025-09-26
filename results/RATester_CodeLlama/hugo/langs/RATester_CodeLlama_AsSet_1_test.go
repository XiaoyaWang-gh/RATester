package langs

import (
	"fmt"
	"testing"
)

func TestAsSet_1(t *testing.T) {
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

	m := l.AsSet()

	if len(m) != 2 {
		t.Errorf("Expected 2 languages, got %d", len(m))
	}

	if _, ok := m["en"]; !ok {
		t.Errorf("Expected 'en' to be in the set")
	}

	if _, ok := m["no"]; !ok {
		t.Errorf("Expected 'no' to be in the set")
	}
}
