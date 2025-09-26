package customerrors

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFlush_6(t *testing.T) {
	type fields struct {
		code           int
		headerSent     bool
		headerMap      http.Header
		responseWriter http.ResponseWriter
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &codeModifier{
				code:           tt.fields.code,
				headerSent:     tt.fields.headerSent,
				headerMap:      tt.fields.headerMap,
				responseWriter: tt.fields.responseWriter,
			}
			r.Flush()
		})
	}
}
