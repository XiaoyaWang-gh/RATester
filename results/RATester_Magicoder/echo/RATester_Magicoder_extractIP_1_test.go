package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestExtractIP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.RemoteAddr = "192.0.2.1:1234"

	ip := extractIP(req)

	if ip != "192.0.2.1" {
		t.Errorf("Expected IP to be 192.0.2.1, but got %s", ip)
	}
}
