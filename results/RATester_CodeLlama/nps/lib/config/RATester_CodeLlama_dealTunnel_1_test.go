package config

import (
	"fmt"
	"reflect"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestDealTunnel_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *file.Tunnel
	}{
		{
			name: "case1",
			args: args{
				s: "server_port=8080,server_ip=127.0.0.1,mode=tcp,target_addr=127.0.0.1:80,target_ip=127.0.0.1,password=123456,local_path=/tmp,strip_pre=,multi_account=multi_account.json",
			},
			want: &file.Tunnel{
				Ports:        "8080",
				ServerIp:     "127.0.0.1",
				Mode:         "tcp",
				TargetAddr:   "127.0.0.1:80",
				Target:       &file.Target{TargetStr: "127.0.0.1:80"},
				Password:     "123456",
				LocalPath:    "/tmp",
				StripPre:     "",
				MultiAccount: &file.MultiAccount{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := dealTunnel(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dealTunnel() = %v, want %v", got, tt.want)
			}
		})
	}
}
