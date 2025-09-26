package retry

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestServeHTTP_38(t *testing.T) {
	type fields struct {
		attempts        int
		initialInterval time.Duration
		next            http.Handler
		listener        Listener
		name            string
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

			r := &retry{
				attempts:        tt.fields.attempts,
				initialInterval: tt.fields.initialInterval,
				next:            tt.fields.next,
				listener:        tt.fields.listener,
				name:            tt.fields.name,
			}
			r.ServeHTTP(tt.args.rw, tt.args.req)
		})
	}
}
