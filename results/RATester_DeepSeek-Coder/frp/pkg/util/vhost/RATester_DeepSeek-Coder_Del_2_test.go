package vhost

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestDel_2(t *testing.T) {
	r := &Routers{
		indexByDomain: map[string]routerByHTTPUser{
			"example.com": {
				"user1": {{location: "location1"}, {location: "location2"}},
				"user2": {{location: "location3"}, {location: "location4"}},
			},
		},
	}

	testCases := []struct {
		name     string
		domain   string
		location string
		httpUser string
		want     routerByHTTPUser
	}{
		{
			name:     "existing router",
			domain:   "example.com",
			location: "location1",
			httpUser: "user1",
			want: routerByHTTPUser{
				"user1": {{location: "location2"}},
				"user2": {{location: "location3"}, {location: "location4"}},
			},
		},
		{
			name:     "non-existing router",
			domain:   "example.com",
			location: "location5",
			httpUser: "user1",
			want: routerByHTTPUser{
				"user1": {{location: "location1"}, {location: "location2"}},
				"user2": {{location: "location3"}, {location: "location4"}},
			},
		},
		{
			name:     "non-existing domain",
			domain:   "nonexistent.com",
			location: "location1",
			httpUser: "user1",
			want:     map[string][]*Router{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r.Del(tc.domain, tc.location, tc.httpUser)
			assert.Equal(t, tc.want, r.indexByDomain)
		})
	}
}
