package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	json := JSON{}
	w := httptest.NewRecorder()
	json.WriteContentType(w)

	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Expected Content-Type to be application/json, but got %s", w.Header().Get("Content-Type"))
	}
}
