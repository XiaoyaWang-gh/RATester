package middlewares

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestServeHTTP_28(t *testing.T) {
	type fields struct {
		handler *safe.Safe
	}
	type args struct {
		rw  http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			h := &HTTPHandlerSwitcher{
				handler: tt.fields.handler,
			}
			h.ServeHTTP(tt.args.rw, tt.args.req)
		})
	}
}
