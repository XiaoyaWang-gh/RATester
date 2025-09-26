package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	xml := XML{}
	w := httptest.NewRecorder()
	xml.WriteContentType(w)

	if w.Header().Get("Content-Type") != "application/xml" {
		t.Errorf("Expected Content-Type to be 'application/xml', but got '%s'", w.Header().Get("Content-Type"))
	}
}
