package vhost

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExist_1(t *testing.T) {
	r := &Routers{
		indexByDomain: map[string]routerByHTTPUser{
			"example.com": {
				"user1": {
					{
						domain:   "example.com",
						location: "/path1",
						httpUser: "user1",
					},
					{
						domain:   "example.com",
						location: "/path2",
						httpUser: "user1",
					},
				},
			},
		},
	}

	tests := []struct {
		name     string
		host     string
		path     string
		httpUser string
		want     *Router
		want1    bool
	}{
		{
			name:     "Existing route",
			host:     "example.com",
			path:     "/path1",
			httpUser: "user1",
			want: &Router{
				domain:   "example.com",
				location: "/path1",
				httpUser: "user1",
			},
			want1: true,
		},
		{
			name:     "Non-existing route",
			host:     "example.com",
			path:     "/non-existing",
			httpUser: "user1",
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

			got, got1 := r.exist(tt.host, tt.path, tt.httpUser)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Routers.exist() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Routers.exist() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
