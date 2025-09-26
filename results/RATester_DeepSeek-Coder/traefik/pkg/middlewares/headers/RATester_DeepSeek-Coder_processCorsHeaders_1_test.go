package headers

import (
	"fmt"
	"net/http"
	"testing"
)

func TestProcessCorsHeaders_1(t *testing.T) {
	type args struct {
		rw  http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		s    *Header
		args args
		want bool
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

			if got := tt.s.processCorsHeaders(tt.args.rw, tt.args.req); got != tt.want {
				t.Errorf("processCorsHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}
