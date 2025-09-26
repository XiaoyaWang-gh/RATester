package forwardedheaders

import (
	"fmt"
	"testing"
)

func TestRemoveIPv6Zone_1(t *testing.T) {
	type args struct {
		clientIP string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with IPv6 address with zone",
			args: args{clientIP: "fe80::d806:a55d:eb1b:49cc%vEthernet (vmxnet3 Ethernet Adapter - Virtual Switch)"},
			want: "fe80::d806:a55d:eb1b:49cc",
		},
		{
			name: "Test with IPv4 address",
			args: args{clientIP: "192.168.1.1"},
			want: "192.168.1.1",
		},
		{
			name: "Test with empty string",
			args: args{clientIP: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := removeIPv6Zone(tt.args.clientIP); got != tt.want {
				t.Errorf("removeIPv6Zone() = %v, want %v", got, tt.want)
			}
		})
	}
}
