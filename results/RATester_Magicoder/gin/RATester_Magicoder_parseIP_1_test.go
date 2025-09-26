package gin

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestparseIP_1(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want net.IP
	}{
		{
			name: "IPv4",
			ip:   "192.168.1.1",
			want: net.ParseIP("192.168.1.1"),
		},
		{
			name: "IPv6",
			ip:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			want: net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
		},
		{
			name: "Invalid",
			ip:   "invalid",
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

			if got := parseIP(tt.ip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
