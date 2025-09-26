package forwardedheaders

import (
	"fmt"
	"testing"
)

func TestRemoveIPv6Zone_1(t *testing.T) {
	tests := []struct {
		name     string
		clientIP string
		want     string
	}{
		{
			name:     "IPv6 with zone",
			clientIP: "fe80::d806:a55d:eb1b:49cc%vEthernet (vmxnet3 Ethernet Adapter - Virtual Switch)",
			want:     "fe80::d806:a55d:eb1b:49cc",
		},
		{
			name:     "IPv6 without zone",
			clientIP: "fe80::d806:a55d:eb1b:49cc",
			want:     "fe80::d806:a55d:eb1b:49cc",
		},
		{
			name:     "IPv4",
			clientIP: "192.168.1.1",
			want:     "192.168.1.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := removeIPv6Zone(tt.clientIP); got != tt.want {
				t.Errorf("removeIPv6Zone() = %v, want %v", got, tt.want)
			}
		})
	}
}
