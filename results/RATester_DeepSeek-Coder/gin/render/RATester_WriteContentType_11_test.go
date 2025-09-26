package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := httptest.NewRecorder()
	r := ProtoBuf{}
	r.WriteContentType(w)

	result := w.Header().Get("Content-Type")
	expected := "application/protobuf"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
