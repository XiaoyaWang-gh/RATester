package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Reader{
		ContentType: "text/html",
	}
	w := httptest.NewRecorder()
	r.WriteContentType(w)
	if w.Header().Get("Content-Type") != "text/html" {
		t.Errorf("Content-Type header is not set correctly")
	}
}
