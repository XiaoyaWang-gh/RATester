package auth

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"
)

func TestServeHTTP_17(t *testing.T) {
	type fields struct {
		address                  string
		authResponseHeaders      []string
		authResponseHeadersRegex *regexp.Regexp
		next                     http.Handler
		name                     string
		client                   http.Client
		trustForwardHeader       bool
		authRequestHeaders       []string
		addAuthCookiesToResponse map[string]struct{}
		headerField              string
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

			fa := &forwardAuth{
				address:                  tt.fields.address,
				authResponseHeaders:      tt.fields.authResponseHeaders,
				authResponseHeadersRegex: tt.fields.authResponseHeadersRegex,
				next:                     tt.fields.next,
				name:                     tt.fields.name,
				client:                   tt.fields.client,
				trustForwardHeader:       tt.fields.trustForwardHeader,
				authRequestHeaders:       tt.fields.authRequestHeaders,
				addAuthCookiesToResponse: tt.fields.addAuthCookiesToResponse,
				headerField:              tt.fields.headerField,
			}
			fa.ServeHTTP(tt.args.rw, tt.args.req)
		})
	}
}
