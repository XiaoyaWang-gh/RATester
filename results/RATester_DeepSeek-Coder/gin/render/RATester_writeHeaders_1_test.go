package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Reader{
		ContentType: "application/json",
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Custom-Header": "Custom Value",
		},
	}

	w := httptest.NewRecorder()
	r.writeHeaders(w, r.Headers)

	header := w.Result().Header

	// Test if Content-Type header is set correctly
	if header.Get("Content-Type") != "application/json" {
		t.Errorf("Expected Content-Type to be 'application/json', got '%s'", header.Get("Content-Type"))
	}

	// Test if Custom-Header is set correctly
	if header.Get("Custom-Header") != "Custom Value" {
		t.Errorf("Expected Custom-Header to be 'Custom Value', got '%s'", header.Get("Custom-Header"))
	}
}
