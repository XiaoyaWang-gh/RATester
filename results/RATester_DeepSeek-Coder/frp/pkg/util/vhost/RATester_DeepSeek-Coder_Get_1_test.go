package vhost

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_1(t *testing.T) {
	type args struct {
		host     string
		path     string
		httpUser string
	}
	tests := []struct {
		name      string
		r         *Routers
		args      args
		wantVr    *Router
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

			gotVr, gotExist := tt.r.Get(tt.args.host, tt.args.path, tt.args.httpUser)
			if !reflect.DeepEqual(gotVr, tt.wantVr) {
				t.Errorf("Routers.Get() gotVr = %v, want %v", gotVr, tt.wantVr)
			}
			if gotExist != tt.wantExist {
				t.Errorf("Routers.Get() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}
