package vhost

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExist_1(t *testing.T) {
	type args struct {
		host     string
		path     string
		httpUser string
	}
	tests := []struct {
		name      string
		r         *Routers
		args      args
		wantRoute *Router
		wantExist bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotRoute, gotExist := tt.r.exist(tt.args.host, tt.args.path, tt.args.httpUser)
			if !reflect.DeepEqual(gotRoute, tt.wantRoute) {
				t.Errorf("Routers.exist() gotRoute = %v, want %v", gotRoute, tt.wantRoute)
			}
			if gotExist != tt.wantExist {
				t.Errorf("Routers.exist() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}
