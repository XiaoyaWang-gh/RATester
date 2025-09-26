package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := httptest.NewRecorder()
	r := AsciiJSON{}
	r.WriteContentType(w)

	if w.Header().Get("Content-Type") != "application/json; charset=utf-8" {
		t.Errorf("Expected 'application/json; charset=utf-8', got '%s'", w.Header().Get("Content-Type"))
	}
}
