package middlewares

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"testing"
)

func TestHijack_2(t *testing.T) {
	type fields struct {
		req         *http.Request
		rw          http.ResponseWriter
		headersSent bool
		code        int
		modifier    func(*http.Response) error
		modified    bool
		modifierErr error
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

			r := &ResponseModifier{
				req:         tt.fields.req,
				rw:          tt.fields.rw,
				headersSent: tt.fields.headersSent,
				code:        tt.fields.code,
				modifier:    tt.fields.modifier,
				modified:    tt.fields.modified,
				modifierErr: tt.fields.modifierErr,
			}
			got, got1, err := r.Hijack()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResponseModifier.Hijack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResponseModifier.Hijack() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ResponseModifier.Hijack() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
