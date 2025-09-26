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
		ContentType: "application/json",
	}

	w := httptest.NewRecorder()
	r.WriteContentType(w)

	resp := w.Result()
	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Expected Content-Type header to be 'application/json', but got '%s'", resp.Header.Get("Content-Type"))
	}
}
