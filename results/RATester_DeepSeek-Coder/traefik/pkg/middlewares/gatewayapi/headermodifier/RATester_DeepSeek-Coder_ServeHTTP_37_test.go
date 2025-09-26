package headermodifier

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeHTTP_37(t *testing.T) {
	type args struct {
		rw  http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		r    *responseHeaderModifier
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

			r := &responseHeaderModifier{
				next:   tt.r.next,
				name:   tt.r.name,
				set:    tt.r.set,
				add:    tt.r.add,
				remove: tt.r.remove,
			}
			r.ServeHTTP(tt.args.rw, tt.args.req)
		})
	}
}
