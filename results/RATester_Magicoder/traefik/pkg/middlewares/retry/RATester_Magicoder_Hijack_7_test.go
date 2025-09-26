package retry

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"testing"
)

func TestHijack_7(t *testing.T) {
	type fields struct {
		responseWriter http.ResponseWriter
		headers        http.Header
		shouldRetry    bool
		written        bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    net.Conn
		want1   *bufio.ReadWriter
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
				responseWriter: tt.fields.responseWriter,
				headers:        tt.fields.headers,
				shouldRetry:    tt.fields.shouldRetry,
				written:        tt.fields.written,
			}
			got, got1, err := r.Hijack()
			if (err != nil) != tt.wantErr {
				t.Errorf("responseWriter.Hijack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("responseWriter.Hijack() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("responseWriter.Hijack() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
