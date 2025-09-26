package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRender_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Reader{
		ContentType:   "text/html",
		ContentLength: 100,
		Reader:        strings.NewReader("hello world"),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}
	w := httptest.NewRecorder()
	err := r.Render(w)
	if err != nil {
		t.Fatal(err)
	}
	if w.Code != http.StatusOK {
		t.Errorf("w.Code = %d, want %d", w.Code, http.StatusOK)
	}
	if w.Header().Get("Content-Type") != "text/html" {
		t.Errorf("w.Header().Get(\"Content-Type\") = %q, want %q", w.Header().Get("Content-Type"), "text/html")
	}
	if w.Header().Get("Content-Length") != "100" {
		t.Errorf("w.Header().Get(\"Content-Length\") = %q, want %q", w.Header().Get("Content-Length"), "100")
	}
	if w.Body.String() != "hello world" {
		t.Errorf("w.Body.String() = %q, want %q", w.Body.String(), "hello world")
	}
}
