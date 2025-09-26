package service

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestBuildSingleHostProxy_1(t *testing.T) {
	type args struct {
		target         *url.URL
		passHostHeader bool
		flushInterval  time.Duration
		roundTripper   http.RoundTripper
		bufferPool     httputil.BufferPool
	}
	tests := []struct {
		name string
		args args
		want http.Handler
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

			if got := buildSingleHostProxy(tt.args.target, tt.args.passHostHeader, tt.args.flushInterval, tt.args.roundTripper, tt.args.bufferPool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildSingleHostProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
