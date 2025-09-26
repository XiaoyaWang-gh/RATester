package urls

import (
	"fmt"
	"net/url"
	"testing"
)

func TestURL_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := BaseURL{
		url: &url.URL{
			Scheme: "https",
			Host:   "example.com",
			Path:   "/path",
		},
	}

	if b.URL().String() != "https://example.com/path" {
		t.Errorf("expected %s, got %s", "https://example.com/path", b.URL().String())
	}
}
