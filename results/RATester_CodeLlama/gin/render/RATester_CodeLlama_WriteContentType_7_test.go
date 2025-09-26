package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var r String
	r.Format = "plain"
	r.Data = []any{"hello"}
	w := httptest.NewRecorder()
	r.WriteContentType(w)
	if w.Header().Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Errorf("Content-Type is not text/plain; charset=utf-8")
	}
}
