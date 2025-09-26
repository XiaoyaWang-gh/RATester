package helpers

import (
	"fmt"
	"testing"
)

func TestURLEscape_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &PathSpec{}
	uri := "https://example.com/path/to/file.html"
	expected := "https://example.com/path/to/file.html"
	actual := p.URLEscape(uri)
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
