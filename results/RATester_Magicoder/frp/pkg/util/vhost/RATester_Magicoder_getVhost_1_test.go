package vhost

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetVhost_1(t *testing.T) {
	rp := &HTTPReverseProxy{
		vhostRouter: &Routers{
			indexByDomain: map[string]routerByHTTPUser{
				"example.com": {
					"": {
						&Router{
							domain:   "example.com",
							location: "/",
							httpUser: "user1",
							payload:  "payload1",
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
		want     *Router
		want1    bool
	}{
		{
			name:     "Existing router",
			domain:   "example.com",
			location: "/",
			httpUser: "user1",
			want: &Router{
				domain:   "example.com",
				location: "/",
				httpUser: "user1",
				payload:  "payload1",
			},
			want1: true,
		},
		{
			name:     "Non-existing router",
			domain:   "example.com",
			location: "/",
			httpUser: "user2",
			want:     nil,
			want1:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := rp.getVhost(tt.domain, tt.location, tt.httpUser)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getVhost() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getVhost() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
