package proxy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHttpConnectListen_1(t *testing.T) {
	type args struct {
		domain          string
		routeByHTTPUser string
		httpUser        string
		httpPwd         string
		addrs           []string
	}
	tests := []struct {
		name    string
		pxy     *TCPMuxProxy
		args    args
		want    []string
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

			got, err := tt.pxy.httpConnectListen(tt.args.domain, tt.args.routeByHTTPUser, tt.args.httpUser, tt.args.httpPwd, tt.args.addrs)
			if (err != nil) != tt.wantErr {
				t.Errorf("TCPMuxProxy.httpConnectListen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TCPMuxProxy.httpConnectListen() = %v, want %v", got, tt.want)
			}
		})
	}
}
