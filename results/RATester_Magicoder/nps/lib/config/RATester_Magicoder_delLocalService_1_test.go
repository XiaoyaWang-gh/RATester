package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestdelLocalService_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *LocalServer
	}{
		{
			name: "Test case 1",
			args: args{s: "local_port=8080&local_ip=127.0.0.1&password=password123&target_addr=192.168.1.1"},
			want: &LocalServer{
				Port:     8080,
				Ip:       "127.0.0.1",
				Password: "password123",
				Target:   "192.168.1.1",
			},
		},
		{
			name: "Test case 2",
			args: args{s: "local_port=8081&local_ip=127.0.0.2&password=password1234&target_addr=192.168.1.2"},
			want: &LocalServer{
				Port:     8081,
				Ip:       "127.0.0.2",
				Password: "password1234",
				Target:   "192.168.1.2",
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := delLocalService(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delLocalService() = %v, want %v", got, tt.want)
			}
		})
	}
}
