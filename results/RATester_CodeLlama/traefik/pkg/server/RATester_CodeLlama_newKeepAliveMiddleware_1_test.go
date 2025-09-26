package server

import (
	"fmt"
	"net/http"
	"testing"
	"time"

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
	}{
		{
			name: "test",
			args: args{
				next:        http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {}),
				maxRequests: 10,
				maxTime:     ptypes.Duration(time.Second * 10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			newKeepAliveMiddleware(tt.args.next, tt.args.maxRequests, tt.args.maxTime)
		})
	}
}
