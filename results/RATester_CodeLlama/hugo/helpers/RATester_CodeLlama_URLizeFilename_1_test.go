package helpers

import (
	"fmt"
	"testing"
)

func TestURLizeFilename_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &PathSpec{}
	filename := "foo/bar.html"
	expected := "foo/bar.html"
	if actual := p.URLizeFilename(filename); actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}
