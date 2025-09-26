package vhost

import (
	"fmt"
	"testing"
)

func TestDel_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Routers{
		indexByDomain: map[string]routerByHTTPUser{
			"example.com": {
				"user1": {
					{location: "location1"},
					{location: "location2"},
				},
			},
		},
	}

	r.Del("example.com", "location1", "user1")

	if len(r.indexByDomain["example.com"]["user1"]) != 1 {
		t.Errorf("Expected 1 router, got %d", len(r.indexByDomain["example.com"]["user1"]))
	}

	if r.indexByDomain["example.com"]["user1"][0].location != "location2" {
		t.Errorf("Expected location2, got %s", r.indexByDomain["example.com"]["user1"][0].location)
	}
}
