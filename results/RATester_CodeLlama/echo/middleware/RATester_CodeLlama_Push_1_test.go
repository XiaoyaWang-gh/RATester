package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := new(gzipResponseWriter)
	w.ResponseWriter = httptest.NewRecorder()
	w.buffer = new(bytes.Buffer)
	w.minLength = 10
	w.minLengthExceeded = true
	w.code = 200
	target := "https://example.com"
	opts := &http.PushOptions{
		Method: "GET",
		Header: http.Header{
			"X-Test": []string{"foo"},
		},
	}
	err := w.Push(target, opts)
	if err != http.ErrNotSupported {
		t.Errorf("Push = %v; want http.ErrNotSupported", err)
	}
}
