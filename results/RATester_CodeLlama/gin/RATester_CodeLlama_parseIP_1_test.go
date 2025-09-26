package gin

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestParseIP_1(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want net.IP
	}{
		{
			name: "test case 1",
			args: args{
				ip: "192.168.1.1",
			},
			want: net.IPv4(192, 168, 1, 1),
		},
		{
			name: "test case 2",
			args: args{
				ip: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			},
			want: net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := parseIP(tt.args.ip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
