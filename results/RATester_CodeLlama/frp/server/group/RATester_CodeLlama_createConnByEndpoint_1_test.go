package group

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestCreateConnByEndpoint_1(t *testing.T) {
	type args struct {
		endpoint   string
		remoteAddr string
	}
	tests := []struct {
		name    string
		args    args
		want    net.Conn
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				endpoint:   "endpoint",
				remoteAddr: "remoteAddr",
			},
			want:    nil,
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

			g := &HTTPGroup{
				group:           "group",
				groupKey:        "groupKey",
				domain:          "domain",
				location:        "location",
				routeByHTTPUser: "routeByHTTPUser",
			}
			got, err := g.createConnByEndpoint(tt.args.endpoint, tt.args.remoteAddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("createConnByEndpoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createConnByEndpoint() got = %v, want %v", got, tt.want)
			}
		})
	}
}
