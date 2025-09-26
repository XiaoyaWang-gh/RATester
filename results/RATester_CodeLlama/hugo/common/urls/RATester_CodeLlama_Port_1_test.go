package urls

import (
	"fmt"
	"net/url"
	"testing"
)

func TestPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := BaseURL{url: &url.URL{Scheme: "http", Host: "example.com:8080"}}
	if b.Port() != 8080 {
		t.Errorf("Port() = %d, want 8080", b.Port())
	}
}
