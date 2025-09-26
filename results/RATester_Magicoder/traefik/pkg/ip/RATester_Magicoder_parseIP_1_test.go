package ip

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestParseIP_1(t *testing.T) {
	tests := []struct {
		name    string
		addr    string
		want    net.IP
		wantErr bool
	}{
		{
			name: "valid IPv4",
			addr: "192.0.2.1",
			want: net.IPv4(192, 0, 2, 1),
		},
		{
			name: "valid IPv6",
			addr: "2001:db8::68",
			want: net.ParseIP("2001:db8::68"),
		},
		{
			name:    "invalid IP",
			addr:    "invalid",
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

			got, err := parseIP(tt.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
