package vhost

import (
	"fmt"
	"testing"
)

func TestUnRegister_1(t *testing.T) {
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
		Location:        "/api",
		RouteByHTTPUser: "user1",
	}

	rp.vhostRouter.Add(routeCfg.Domain, routeCfg.Location, routeCfg.RouteByHTTPUser, nil)

	rp.UnRegister(routeCfg)

	_, exist := rp.vhostRouter.Get(routeCfg.Domain, routeCfg.Location, routeCfg.RouteByHTTPUser)
	if exist {
		t.Errorf("Expected route to be removed, but it still exists")
	}
}
