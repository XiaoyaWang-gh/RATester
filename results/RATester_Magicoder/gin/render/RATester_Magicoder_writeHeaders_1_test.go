package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestwriteHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Reader{
		ContentType: "text/plain",
		Headers: map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": "100",
		},
	}

	w := httptest.NewRecorder()
	r.writeHeaders(w, r.Headers)

	header := w.Header()

	if header.Get("Content-Type") != "text/plain" {
		t.Errorf("Expected Content-Type to be 'text/plain', but got '%s'", header.Get("Content-Type"))
	}

	if header.Get("Content-Length") != "100" {
		t.Errorf("Expected Content-Length to be '100', but got '%s'", header.Get("Content-Length"))
	}
}
