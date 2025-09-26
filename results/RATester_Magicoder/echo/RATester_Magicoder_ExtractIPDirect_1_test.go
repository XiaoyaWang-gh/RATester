package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestExtractIPDirect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	extractor := ExtractIPDirect()
	ip := extractor(req)

	if ip != "" {
		t.Errorf("Expected empty string, got %s", ip)
	}
}
