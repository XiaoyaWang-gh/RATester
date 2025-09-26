package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := httptest.NewRecorder()
	value := []string{"application/json; charset=utf-8"}
	writeContentType(w, value)

	header := w.Header()
	if val := header["Content-Type"]; len(val) != 1 || val[0] != "application/json; charset=utf-8" {
		t.Errorf("Expected Content-Type to be 'application/json; charset=utf-8', got %v", val)
	}
}
