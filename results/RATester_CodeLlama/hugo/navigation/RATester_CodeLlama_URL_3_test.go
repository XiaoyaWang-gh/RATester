package navigation

import (
	"fmt"
	"testing"
)

func TestURL_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &MenuEntry{
		ConfiguredURL: "https://example.com",
	}

	if m.URL() != "https://example.com" {
		t.Errorf("expected %q, got %q", "https://example.com", m.URL())
	}
}
