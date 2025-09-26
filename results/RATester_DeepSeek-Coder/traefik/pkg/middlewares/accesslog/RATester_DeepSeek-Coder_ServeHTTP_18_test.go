package accesslog

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeHTTP_18(t *testing.T) {
	type args struct {
		rw   http.ResponseWriter
		req  *http.Request
		next http.Handler
	}
	tests := []struct {
		name string
		h    *Handler
		args args
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

			h := &Handler{}
			h.ServeHTTP(tt.args.rw, tt.args.req, tt.args.next)
		})
	}
}
