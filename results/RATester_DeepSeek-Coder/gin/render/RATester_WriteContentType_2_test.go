package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := httptest.NewRecorder()
	r := IndentedJSON{}
	r.WriteContentType(w)

	result := w.Header().Get("Content-Type")
	expected := "application/json"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
