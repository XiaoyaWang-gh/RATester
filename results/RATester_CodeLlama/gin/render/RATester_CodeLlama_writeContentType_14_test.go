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
	writeContentType(w, []string{"application/json"})
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type is not application/json")
	}
}
