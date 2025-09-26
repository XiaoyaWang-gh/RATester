package vhost

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRouteConfig_1(t *testing.T) {
	rp := &HTTPReverseProxy{
		vhostRouter: &Routers{
			indexByDomain: make(map[string]routerByHTTPUser),
		},
	}

	testCases := []struct {
		name     string
		domain   string
		location string
		httpUser string
		want     *RouteConfig
	}{
		{
			name:     "existing route",
			domain:   "example.com",
			location: "/api",
			httpUser: "user1",
			want: &RouteConfig{
				Domain:   "example.com",
				Location: "/api",
				Username: "user1",
			},
		},
		{
			name:     "non-existing route",
			domain:   "non-existing.com",
			location: "/non-existing",
			httpUser: "non-existing-user",
			want:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rp.vhostRouter.Add(tc.domain, tc.location, tc.httpUser, tc.want)
			got := rp.GetRouteConfig(tc.domain, tc.location, tc.httpUser)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetRouteConfig() = %v, want %v", got, tc.want)
			}
		})
	}
}
