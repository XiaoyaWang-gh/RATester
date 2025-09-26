package urls

import (
	"fmt"
	"net/url"
	"testing"
)

func TestPath_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := BaseURL{
		url: &url.URL{
			Path: "/path/to/file",
		},
	}
	if b.Path() != "/path/to/file" {
		t.Errorf("Expected %s, got %s", "/path/to/file", b.Path())
	}
}
