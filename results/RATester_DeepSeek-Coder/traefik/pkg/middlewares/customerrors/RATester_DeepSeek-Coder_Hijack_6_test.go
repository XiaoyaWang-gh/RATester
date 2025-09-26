package customerrors

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"testing"
)

func TestHijack_6(t *testing.T) {
	type fields struct {
		code           int
		headerSent     bool
		headerMap      http.Header
		responseWriter http.ResponseWriter
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

			r := &codeModifier{
				code:           tt.fields.code,
				headerSent:     tt.fields.headerSent,
				headerMap:      tt.fields.headerMap,
				responseWriter: tt.fields.responseWriter,
			}
			got, got1, err := r.Hijack()
			if (err != nil) != tt.wantErr {
				t.Errorf("codeModifier.Hijack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("codeModifier.Hijack() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("codeModifier.Hijack() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
