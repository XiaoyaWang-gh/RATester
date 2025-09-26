package server

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	ptypes "github.com/traefik/paerser/types"
)

func TestNewKeepAliveMiddleware_1(t *testing.T) {
	type args struct {
		next        http.Handler
		maxRequests int
		maxTime     ptypes.Duration
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

			if got := newKeepAliveMiddleware(tt.args.next, tt.args.maxRequests, tt.args.maxTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newKeepAliveMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
