package server

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetConfByType_1(t *testing.T) {
	type args struct {
		proxyType string
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Test TCP Proxy",
			args: args{proxyType: "TCP"},
			want: &TCPOutConf{},
		},
		{
			name: "Test TCPMUX Proxy",
			args: args{proxyType: "TCPMUX"},
			want: &TCPMuxOutConf{},
		},
		{
			name: "Test UDP Proxy",
			args: args{proxyType: "UDP"},
			want: &UDPOutConf{},
		},
		{
			name: "Test HTTP Proxy",
			args: args{proxyType: "HTTP"},
			want: &HTTPOutConf{},
		},
		{
			name: "Test HTTPS Proxy",
			args: args{proxyType: "HTTPS"},
			want: &HTTPSOutConf{},
		},
		{
			name: "Test STCP Proxy",
			args: args{proxyType: "STCP"},
			want: &STCPOutConf{},
		},
		{
			name: "Test XTCP Proxy",
			args: args{proxyType: "XTCP"},
			want: &XTCPOutConf{},
		},
		{
			name: "Test Unknown Proxy",
			args: args{proxyType: "UNKNOWN"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getConfByType(tt.args.proxyType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfByType() = %v, want %v", got, tt.want)
			}
		})
	}
}
