package compress

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClose_2(t *testing.T) {
	type fields struct {
		rw                   http.ResponseWriter
		compressionWriter    *compressionWriter
		minSize              int
		excludedContentTypes []parsedContentType
		includedContentTypes []parsedContentType
		buf                  []byte
		hijacked             bool
		compressionStarted   bool
		compressionDisabled  bool
		headersSent          bool
		seenData             bool
		statusCodeSet        bool
		statusCode           int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
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

			r := &responseWriter{
				rw:                   tt.fields.rw,
				compressionWriter:    tt.fields.compressionWriter,
				minSize:              tt.fields.minSize,
				excludedContentTypes: tt.fields.excludedContentTypes,
				includedContentTypes: tt.fields.includedContentTypes,
				buf:                  tt.fields.buf,
				hijacked:             tt.fields.hijacked,
				compressionStarted:   tt.fields.compressionStarted,
				compressionDisabled:  tt.fields.compressionDisabled,
				headersSent:          tt.fields.headersSent,
				seenData:             tt.fields.seenData,
				statusCodeSet:        tt.fields.statusCodeSet,
				statusCode:           tt.fields.statusCode,
			}
			if err := r.close(); (err != nil) != tt.wantErr {
				t.Errorf("responseWriter.close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
