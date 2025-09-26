package stripprefix

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeRequest_1(t *testing.T) {
	type args struct {
		rw     http.ResponseWriter
		req    *http.Request
		prefix string
	}
	tests := []struct {
		name string
		s    *stripPrefix
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

			s := &stripPrefix{
				next:       tt.s.next,
				prefixes:   tt.s.prefixes,
				name:       tt.s.name,
				forceSlash: tt.s.forceSlash,
			}
			s.serveRequest(tt.args.rw, tt.args.req, tt.args.prefix)
		})
	}
}
