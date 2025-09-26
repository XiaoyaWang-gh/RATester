package context

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestHijack_1(t *testing.T) {
	type fields struct {
		ResponseWriter http.ResponseWriter
		Started        bool
		Status         int
		Elapsed        time.Duration
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

			r := &Response{
				ResponseWriter: tt.fields.ResponseWriter,
				Started:        tt.fields.Started,
				Status:         tt.fields.Status,
				Elapsed:        tt.fields.Elapsed,
			}
			got, got1, err := r.Hijack()
			if (err != nil) != tt.wantErr {
				t.Errorf("Response.Hijack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.Hijack() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Response.Hijack() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
