package urls

import (
	"fmt"
	"testing"
)

func TestString_23(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := BaseURL{
		WithPath: "https://example.com/path",
	}
	if b.String() != "https://example.com/path" {
		t.Errorf("BaseURL.String() = %q, want %q", b.String(), "https://example.com/path")
	}
}
