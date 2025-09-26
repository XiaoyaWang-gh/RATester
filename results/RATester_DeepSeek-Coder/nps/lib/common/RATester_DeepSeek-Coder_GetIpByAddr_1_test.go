package common

import (
	"fmt"
	"testing"
)

func TestGetIpByAddr_1(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{addr: "192.168.1.1:8080"},
			want: "192.168.1.1",
		},
		{
			name: "Test case 2",
			args: args{addr: "10.0.0.1:80"},
			want: "10.0.0.1",
		},
		{
			name: "Test case 3",
			args: args{addr: "172.16.0.1:22"},
			want: "172.16.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetIpByAddr(tt.args.addr); got != tt.want {
				t.Errorf("GetIpByAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
