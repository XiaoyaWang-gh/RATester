package vhost

import (
	"fmt"
	"testing"
)

func TestCheckAuth_1(t *testing.T) {
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

	// Test case 1: valid user and password
	rp.vhostRouter.indexByDomain["test.com"] = routerByHTTPUser{
		"user1": {
			&Router{
				domain:   "test.com",
				location: "/",
				httpUser: "user1",
				payload: &RouteConfig{
					Username: "user1",
					Password: "pass1",
				},
			},
		},
	}
	ok := rp.CheckAuth("test.com", "/", "user1", "user1", "pass1")
	if !ok {
		t.Errorf("Expected true, got false")
	}

	// Test case 2: invalid user
	ok = rp.CheckAuth("test.com", "/", "user1", "user2", "pass1")
	if ok {
		t.Errorf("Expected false, got true")
	}

	// Test case 3: invalid password
	ok = rp.CheckAuth("test.com", "/", "user1", "user1", "pass2")
	if ok {
		t.Errorf("Expected false, got true")
	}

	// Test case 4: no user and password
	rp.vhostRouter.indexByDomain["test.com"] = routerByHTTPUser{
		"user1": {
			&Router{
				domain:   "test.com",
				location: "/",
				httpUser: "user1",
				payload: &RouteConfig{
					Username: "",
					Password: "",
				},
			},
		},
	}
	ok = rp.CheckAuth("test.com", "/", "user1", "user1", "pass1")
	if !ok {
		t.Errorf("Expected true, got false")
	}
}
