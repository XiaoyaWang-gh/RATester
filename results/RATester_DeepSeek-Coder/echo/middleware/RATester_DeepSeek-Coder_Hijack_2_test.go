package middleware

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"reflect"
	"testing"
)

func TestHijack_2(t *testing.T) {
	type fields struct {
		Writer         io.Writer
		ResponseWriter http.ResponseWriter
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

			w := &bodyDumpResponseWriter{
				Writer:         tt.fields.Writer,
				ResponseWriter: tt.fields.ResponseWriter,
			}
			got, got1, err := w.Hijack()
			if (err != nil) != tt.wantErr {
				t.Errorf("bodyDumpResponseWriter.Hijack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bodyDumpResponseWriter.Hijack() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("bodyDumpResponseWriter.Hijack() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
