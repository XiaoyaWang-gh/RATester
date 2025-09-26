package metrics

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRetried_1(t *testing.T) {
	type fields struct {
		retryMetrics retryMetrics
		serviceName  string
	}
	type args struct {
		req *http.Request
		n   int
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

			m := &RetryListener{
				retryMetrics: tt.fields.retryMetrics,
				serviceName:  tt.fields.serviceName,
			}
			m.Retried(tt.args.req, tt.args.n)
		})
	}
}
