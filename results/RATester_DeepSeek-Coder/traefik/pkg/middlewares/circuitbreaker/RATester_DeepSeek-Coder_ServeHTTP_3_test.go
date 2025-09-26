package circuitbreaker

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/vulcand/oxy/v2/cbreaker"
)

func TestServeHTTP_3(t *testing.T) {
	type fields struct {
		circuitBreaker *cbreaker.CircuitBreaker
		name           string
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

			c := &circuitBreaker{
				circuitBreaker: tt.fields.circuitBreaker,
				name:           tt.fields.name,
			}
			c.ServeHTTP(tt.args.rw, tt.args.req)
		})
	}
}
