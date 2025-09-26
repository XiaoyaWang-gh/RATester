package accesslog

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRetried_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	saveRetries := SaveRetries{}
	saveRetries.Retried(req, 3)

	if req.Header.Get("Retry-Attempts") != "2" {
		t.Errorf("Expected Retry-Attempts header to be 2, but got %s", req.Header.Get("Retry-Attempts"))
	}
}
