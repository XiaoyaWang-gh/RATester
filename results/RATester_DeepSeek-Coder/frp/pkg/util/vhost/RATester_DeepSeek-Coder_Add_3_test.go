package vhost

import (
	"fmt"
	"testing"
)

func TestAdd_3(t *testing.T) {
	type args struct {
		domain   string
		location string
		httpUser string
		payload  interface{}
	}
	tests := []struct {
		name    string
		r       *Routers
		args    args
		wantErr bool
	}{
		{
			name: "TestAdd_0",
			r: &Routers{
				indexByDomain: make(map[string]routerByHTTPUser),
			},
			args: args{
				domain:   "example.com",
				location: "/",
				httpUser: "user1",
				payload:  "payload1",
			},
			wantErr: false,
		},
		{
			name: "TestAdd_1",
			r: &Routers{
				indexByDomain: make(map[string]routerByHTTPUser),
			},
			args: args{
				domain:   "example.com",
				location: "/",
				httpUser: "user1",
				payload:  "payload1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &Routers{
				indexByDomain: tt.r.indexByDomain,
			}
			if err := r.Add(tt.args.domain, tt.args.location, tt.args.httpUser, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("Routers.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
