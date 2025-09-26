package buffering

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeHTTP_39(t *testing.T) {
	type args struct {
		rw  http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		b    *buffer
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

			tt.b.ServeHTTP(tt.args.rw, tt.args.req)
		})
	}
}
