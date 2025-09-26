package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := httptest.NewRecorder()
	msgpack := MsgPack{}
	msgpack.WriteContentType(w)

	result := w.Header().Get("Content-Type")
	expected := "application/msgpack"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
