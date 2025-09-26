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

	msgpack := MsgPack{}
	w := httptest.NewRecorder()
	msgpack.WriteContentType(w)

	if w.Header().Get("Content-Type") != "application/msgpack" {
		t.Errorf("Expected Content-Type to be 'application/msgpack', but got '%s'", w.Header().Get("Content-Type"))
	}
}
