package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := httptest.NewRecorder()
	y := YAML{Data: "test"}

	y.WriteContentType(w)

	expectedContentType := "application/x-yaml; charset=utf-8"
	if w.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("Expected Content-Type to be %s, got %s", expectedContentType, w.Header().Get("Content-Type"))
	}
}
