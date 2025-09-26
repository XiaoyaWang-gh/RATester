package capture

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"testing"
)

func TestHijack_4(t *testing.T) {
	type fields struct {
		rw     http.ResponseWriter
		status int
		size   int64
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

			crw := &captureResponseWriter{
				rw:     tt.fields.rw,
				status: tt.fields.status,
				size:   tt.fields.size,
			}
			got, got1, err := crw.Hijack()
			if (err != nil) != tt.wantErr {
				t.Errorf("captureResponseWriter.Hijack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("captureResponseWriter.Hijack() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("captureResponseWriter.Hijack() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
