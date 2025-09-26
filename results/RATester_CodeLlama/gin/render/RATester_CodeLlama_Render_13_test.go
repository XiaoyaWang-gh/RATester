package render

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	data := Data{
		ContentType: "text/html",
		Data:        []byte("<html><body>Hello, world!</body></html>"),
	}
	w := httptest.NewRecorder()
	// when
	err := data.Render(w)
	// then
	if err != nil {
		t.Errorf("Render() error = %v", err)
		return
	}
	if w.Code != http.StatusOK {
		t.Errorf("Render() code = %v, want %v", w.Code, http.StatusOK)
	}
	if w.Header().Get("Content-Type") != data.ContentType {
		t.Errorf("Render() content-type = %v, want %v", w.Header().Get("Content-Type"), data.ContentType)
	}
	if !bytes.Equal(w.Body.Bytes(), data.Data) {
		t.Errorf("Render() body = %v, want %v", w.Body.Bytes(), data.Data)
	}
}
