package config

import (
	"fmt"
	"reflect"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestDealHost_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *file.Host
	}{
		{
			name: "case1",
			args: args{
				s: "host=127.0.0.1,target_addr=127.0.0.1:8080,host_change=127.0.0.1,scheme=http,location=/,header_host:127.0.0.1,header_x-forwarded-for:127.0.0.1",
			},
			want: &file.Host{
				Host:         "127.0.0.1",
				HeaderChange: "host:127.0.0.1\nx-forwarded-for:127.0.0.1",
				HostChange:   "127.0.0.1",
				Location:     "/",
				Scheme:       "http",
				Target: &file.Target{
					TargetStr: "127.0.0.1:8080",
				},
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

			if got := dealHost(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dealHost() = %v, want %v", got, tt.want)
			}
		})
	}
}
