package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := PureJSON{Data: "test"}
	w := httptest.NewRecorder()
	r.WriteContentType(w)

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json; charset=utf-8" {
		t.Errorf("Expected Content-Type to be 'application/json; charset=utf-8', but got '%s'", contentType)
	}
}
