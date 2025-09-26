package wrr

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeHTTP_11(t *testing.T) {
	type fields struct {
		stickyCookie     *stickyCookie
		wantsHealthCheck bool
		handlerMap       map[string]*namedHandler
		handlers         []*namedHandler
		curDeadline      float64
		status           map[string]struct{}
		updaters         []func(bool)
	}
	type args struct {
		w   http.ResponseWriter
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

			b := &Balancer{
				stickyCookie:     tt.fields.stickyCookie,
				wantsHealthCheck: tt.fields.wantsHealthCheck,
				handlerMap:       tt.fields.handlerMap,
				handlers:         tt.fields.handlers,
				curDeadline:      tt.fields.curDeadline,
				status:           tt.fields.status,
				updaters:         tt.fields.updaters,
			}
			b.ServeHTTP(tt.args.w, tt.args.req)
		})
	}
}
