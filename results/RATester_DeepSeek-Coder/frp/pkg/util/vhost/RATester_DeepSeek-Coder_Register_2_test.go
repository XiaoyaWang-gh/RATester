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
		Domain:   "example.com",
		Location: "/api",
	}

	err := rp.Register(routeCfg)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, exist := rp.vhostRouter.indexByDomain[routeCfg.Domain]
	if !exist {
		t.Errorf("Expected domain to be registered, but it was not")
	}
}
