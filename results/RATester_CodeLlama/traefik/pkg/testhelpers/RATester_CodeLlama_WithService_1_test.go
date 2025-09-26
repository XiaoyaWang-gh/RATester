package testhelpers

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestWithService_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "test"
	opts := []func(*dynamic.ServersLoadBalancer){
		func(r *dynamic.ServersLoadBalancer) {
			r.Servers = []dynamic.Server{
				{
					URL: "http://127.0.0.1:80",
				},
			}
		},
	}
	WithService(name, opts...)
}
