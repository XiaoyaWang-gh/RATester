package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	data := Data{
		ContentType: "text/html",
		Data:        []byte("hello world"),
	}
	w := httptest.NewRecorder()
	data.WriteContentType(w)
	if w.Header().Get("Content-Type") != "text/html" {
		t.Errorf("Content-Type is not set")
	}
}
