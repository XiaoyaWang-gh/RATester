package urls

import (
	"fmt"
	"net/url"
	"testing"
)

func TestWithPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := BaseURL{
		url: &url.URL{
			Scheme: "https",
			Host:   "example.com",
		},
	}

	b2, err := b.WithPort(8080)
	if err != nil {
		t.Fatal(err)
	}

	if b2.String() != "https://example.com:8080" {
		t.Errorf("expected %q, got %q", "https://example.com:8080", b2.String())
	}
}
