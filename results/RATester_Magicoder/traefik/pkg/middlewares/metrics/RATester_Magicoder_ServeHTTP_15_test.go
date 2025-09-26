package metrics

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeHTTP_15(t *testing.T) {
	type args struct {
		rw  http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		m    *metricsMiddleware
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

			tt.m.ServeHTTP(tt.args.rw, tt.args.req)
		})
	}
}
