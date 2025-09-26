package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Redirect{
		Code:     301,
		Location: "http://example.com",
	}

	w := httptest.NewRecorder()

	r.WriteContentType(w)

	if w.Header().Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Errorf("Expected Content-Type to be 'text/plain; charset=utf-8', got %s", w.Header().Get("Content-Type"))
	}
}
