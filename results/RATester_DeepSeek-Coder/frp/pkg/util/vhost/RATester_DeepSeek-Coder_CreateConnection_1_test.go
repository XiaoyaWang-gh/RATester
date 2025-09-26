package vhost

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestCreateConnection_1(t *testing.T) {
	type args struct {
		reqRouteInfo *RequestRouteInfo
		byEndpoint   bool
	}
	tests := []struct {
		name    string
		args    args
		want    net.Conn
		wantErr bool
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

			rp := &HTTPReverseProxy{}
			got, err := rp.CreateConnection(tt.args.reqRouteInfo, tt.args.byEndpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPReverseProxy.CreateConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPReverseProxy.CreateConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
