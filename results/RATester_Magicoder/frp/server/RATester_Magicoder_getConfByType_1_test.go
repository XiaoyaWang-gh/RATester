package server

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetConfByType_1(t *testing.T) {
	tests := []struct {
		name      string
		proxyType string
		want      any
	}{
		{
			name:      "TCP",
			proxyType: "TCP",
			want:      &TCPOutConf{},
		},
		{
			name:      "TCPMUX",
			proxyType: "TCPMUX",
			want:      &TCPMuxOutConf{},
		},
		{
			name:      "UDP",
			proxyType: "UDP",
			want:      &UDPOutConf{},
		},
		{
			name:      "HTTP",
			proxyType: "HTTP",
			want:      &HTTPOutConf{},
		},
		{
			name:      "HTTPS",
			proxyType: "HTTPS",
			want:      &HTTPSOutConf{},
		},
		{
			name:      "STCP",
			proxyType: "STCP",
			want:      &STCPOutConf{},
		},
		{
			name:      "XTCP",
			proxyType: "XTCP",
			want:      &XTCPOutConf{},
		},
		{
			name:      "Unknown",
			proxyType: "Unknown",
			want:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getConfByType(tt.proxyType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfByType() = %v, want %v", got, tt.want)
			}
		})
	}
}
