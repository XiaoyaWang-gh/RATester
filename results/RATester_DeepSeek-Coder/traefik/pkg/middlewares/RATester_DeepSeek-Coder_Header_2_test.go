package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeader_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	modifier := &ResponseModifier{
		req: req,
		rw:  rr,
	}

	header := modifier.Header()

	if header == nil {
		t.Errorf("Expected non-nil header, got nil")
	}

	header.Set("Content-Type", "application/json")

	if contentType := header.Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Expected Content-Type to be 'application/json', got %s", contentType)
	}
}
