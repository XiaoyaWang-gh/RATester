package vhost

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRouteConfig_1(t *testing.T) {
	rp := &HTTPReverseProxy{
		vhostRouter: &Routers{
			indexByDomain: map[string]routerByHTTPUser{
				"test.com": {
					"user1": {
						&Router{
							domain:   "test.com",
							location: "/",
							httpUser: "user1",
							payload: &RouteConfig{
								Domain:          "test.com",
								Location:        "/",
								RouteByHTTPUser: "user1",
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name     string
		domain   string
		location string
		httpUser string
		want     *RouteConfig
	}{
		{
			name:     "TestGetRouteConfig_Exist",
			domain:   "test.com",
			location: "/",
			httpUser: "user1",
			want: &RouteConfig{
				Domain:          "test.com",
				Location:        "/",
				RouteByHTTPUser: "user1",
			},
		},
		{
			name:     "TestGetRouteConfig_NotExist",
			domain:   "test.com",
			location: "/",
			httpUser: "user2",
			want:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := rp.GetRouteConfig(tt.domain, tt.location, tt.httpUser)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRouteConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
