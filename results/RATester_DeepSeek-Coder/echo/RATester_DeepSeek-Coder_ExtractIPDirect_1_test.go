package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExtractIPDirect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "192.0.2.0:8080"
	extractor := ExtractIPDirect()
	ip := extractor(req)
	if ip != "192.0.2.0" {
		t.Errorf("Expected IP to be '192.0.2.0', got %s", ip)
	}
}
