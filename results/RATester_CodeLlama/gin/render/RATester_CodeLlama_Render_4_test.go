package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Redirect{
		Code:     http.StatusMovedPermanently,
		Request:  &http.Request{},
		Location: "http://www.example.com",
	}
	w := httptest.NewRecorder()
	err := r.Render(w)
	if err != nil {
		t.Errorf("Render() error = %v", err)
		return
	}
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Render() code = %v, want %v", w.Code, http.StatusMovedPermanently)
	}
	if w.Header().Get("Location") != "http://www.example.com" {
		t.Errorf("Render() location = %v, want %v", w.Header().Get("Location"), "http://www.example.com")
	}
}
