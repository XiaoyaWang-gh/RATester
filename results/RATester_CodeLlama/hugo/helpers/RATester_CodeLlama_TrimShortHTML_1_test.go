package helpers

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/media"
)

func TestTrimShortHTML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ContentSpec{}
	input := []byte("<p>This is a test</p>")
	markup := media.DefaultContentTypes.AsciiDoc.SubType
	expected := []byte("This is a test")
	actual := c.TrimShortHTML(input, markup)
	if !bytes.Equal(expected, actual) {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}
