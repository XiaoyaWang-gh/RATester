package create

import (
	"fmt"
	"net/http"
	"testing"
)

func TestaddDefaultHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	addDefaultHeaders(req)

	if req.Header.Get("User-Agent") != "Hugo Static Site Generator" {
		t.Errorf("Expected User-Agent header to be 'Hugo Static Site Generator', but got '%s'", req.Header.Get("User-Agent"))
	}
}
