package observability

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"testing"
)

func TestHijack_3(t *testing.T) {
	type fields struct {
		ResponseWriter http.ResponseWriter
		status         int
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

			s := &statusCodeRecorder{
				ResponseWriter: tt.fields.ResponseWriter,
				status:         tt.fields.status,
			}
			got, got1, err := s.Hijack()
			if (err != nil) != tt.wantErr {
				t.Errorf("statusCodeRecorder.Hijack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("statusCodeRecorder.Hijack() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("statusCodeRecorder.Hijack() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
