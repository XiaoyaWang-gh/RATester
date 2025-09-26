package vhost

import (
	"fmt"
	"testing"
)

func TestRegister_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rp := &HTTPReverseProxy{
		vhostRouter: &Routers{
			indexByDomain: make(map[string]routerByHTTPUser),
		},
	}

	routeCfg := RouteConfig{
		Domain:          "example.com",
		Location:        "/",
		RouteByHTTPUser: "user",
	}

	err := rp.Register(routeCfg)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, exist := rp.vhostRouter.Get(routeCfg.Domain, routeCfg.Location, routeCfg.RouteByHTTPUser)
	if !exist {
		t.Errorf("Expected route to exist after registration")
	}
}
