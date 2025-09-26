package tcpmux

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestGetHostFromHTTPConnect_1(t *testing.T) {
	type args struct {
		c net.Conn
	}
	tests := []struct {
		name    string
		muxer   *HTTPConnectTCPMuxer
		args    args
		want    net.Conn
		want1   map[string]string
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

			got, got1, err := tt.muxer.getHostFromHTTPConnect(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("getHostFromHTTPConnect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHostFromHTTPConnect() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getHostFromHTTPConnect() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
