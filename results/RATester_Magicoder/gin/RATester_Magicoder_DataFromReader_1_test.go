package gin

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestDataFromReader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		writermem: responseWriter{},
		Request:   &http.Request{},
		Writer:    &responseWriter{},
		Params:    Params{},
	}

	code := 200
	contentLength := int64(10)
	contentType := "text/plain"
	reader := strings.NewReader("test data")
	extraHeaders := map[string]string{"Content-Encoding": "gzip"}

	ctx.DataFromReader(code, contentLength, contentType, reader, extraHeaders)

	// Add assertions here to verify the expected behavior
}
