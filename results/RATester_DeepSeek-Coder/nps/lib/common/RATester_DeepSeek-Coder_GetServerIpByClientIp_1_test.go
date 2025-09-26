package common

import (
	"fmt"
	"net"
	"testing"
)

func TestGetServerIpByClientIp_1(t *testing.T) {
	type args struct {
		clientIp net.IP
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetServerIpByClientIp(tt.args.clientIp); got != tt.want {
				t.Errorf("GetServerIpByClientIp() = %v, want %v", got, tt.want)
			}
		})
	}
}
