package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := httptest.NewRecorder()
	r := JsonpJSON{
		Callback: "callback",
		Data:     "data",
	}
	r.WriteContentType(w)

	expected := "application/json"
	actual := w.Header().Get("Content-Type")
	if actual != expected {
		t.Errorf("Expected Content-Type to be %s, got %s", expected, actual)
	}
}
