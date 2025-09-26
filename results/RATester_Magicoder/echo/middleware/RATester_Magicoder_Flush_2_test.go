package middleware

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestFlush_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &gzipResponseWriter{
		wroteHeader:       false,
		wroteBody:         false,
		minLength:         0,
		minLengthExceeded: false,
		buffer:            bytes.NewBuffer([]byte{}),
		code:              200,
	}

	w.Flush()

	if w.minLengthExceeded != true {
		t.Errorf("Expected minLengthExceeded to be true, but it was false")
	}

	if w.Header().Get(echo.HeaderContentEncoding) != gzipScheme {
		t.Errorf("Expected Header Content-Encoding to be %s, but it was %s", gzipScheme, w.Header().Get(echo.HeaderContentEncoding))
	}

	if w.wroteHeader != true {
		t.Errorf("Expected wroteHeader to be true, but it was false")
	}

	if w.ResponseWriter.Header().Get(echo.HeaderContentEncoding) != gzipScheme {
		t.Errorf("Expected ResponseWriter Header Content-Encoding to be %s, but it was %s", gzipScheme, w.ResponseWriter.Header().Get(echo.HeaderContentEncoding))
	}

	if w.Writer.(*gzip.Writer).Flush() != nil {
		t.Errorf("Expected Writer.Flush() to return nil, but it returned an error")
	}

	if http.NewResponseController(w.ResponseWriter).Flush() != nil {
		t.Errorf("Expected http.NewResponseController(w.ResponseWriter).Flush() to return nil, but it returned an error")
	}
}
