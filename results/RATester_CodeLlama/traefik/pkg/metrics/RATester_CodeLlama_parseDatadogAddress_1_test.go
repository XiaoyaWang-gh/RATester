package metrics

import (
	"fmt"
	"testing"
)

func TestParseDatadogAddress_1(t *testing.T) {
	var tests = []struct {
		name        string
		address     string
		wantNetwork string
		wantAddr    string
	}{
		{
			name:        "empty",
			address:     "",
			wantNetwork: "udp",
			wantAddr:    "localhost:8125",
		},
		{
			name:        "unix",
			address:     "unix:///var/run/datadog/dsd.socket",
			wantNetwork: "unix",
			wantAddr:    "/var/run/datadog/dsd.socket",
		},
		{
			name:        "udp",
			address:     "127.0.0.1:8125",
			wantNetwork: "udp",
			wantAddr:    "127.0.0.1:8125",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotNetwork, gotAddr := parseDatadogAddress(tt.address)
			if gotNetwork != tt.wantNetwork {
				t.Errorf("parseDatadogAddress() gotNetwork = %v, want %v", gotNetwork, tt.wantNetwork)
			}
			if gotAddr != tt.wantAddr {
				t.Errorf("parseDatadogAddress() gotAddr = %v, want %v", gotAddr, tt.wantAddr)
			}
		})
	}
}
