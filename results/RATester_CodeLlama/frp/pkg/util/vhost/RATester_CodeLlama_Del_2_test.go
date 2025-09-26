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
			"domain1": {
				"user1": {
					{
						location: "location1",
						payload:  "payload1",
					},
					{
						location: "location2",
						payload:  "payload2",
					},
				},
			},
		},
	}

	r.Del("domain1", "location1", "user1")

	if len(r.indexByDomain["domain1"]["user1"]) != 1 {
		t.Errorf("expected 1, got %d", len(r.indexByDomain["domain1"]["user1"]))
	}

	if r.indexByDomain["domain1"]["user1"][0].location != "location2" {
		t.Errorf("expected location2, got %s", r.indexByDomain["domain1"]["user1"][0].location)
	}
}
