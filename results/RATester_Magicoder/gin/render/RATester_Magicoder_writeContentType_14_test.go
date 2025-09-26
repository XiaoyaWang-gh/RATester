package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestwriteContentType_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := httptest.NewRecorder()
	value := []string{"application/json"}
	writeContentType(w, value)

	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		t.Errorf("Expected Content-Type header to be set, but it was not.")
	}

	if val := header["Content-Type"][0]; val != "application/json" {
		t.Errorf("Expected Content-Type header to be 'application/json', but got '%s'", val)
	}
}
