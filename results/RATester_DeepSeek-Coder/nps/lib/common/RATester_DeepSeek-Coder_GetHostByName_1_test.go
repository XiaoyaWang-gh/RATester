package common

import (
	"fmt"
	"testing"
)

func TestGetHostByName_1(t *testing.T) {
	type args struct {
		hostname string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with valid domain",
			args: args{hostname: "www.google.com"},
			want: "216.58.217.46",
		},
		{
			name: "Test with invalid domain",
			args: args{hostname: "www.invalid.com"},
			want: "",
		},
		{
			name: "Test with IP address",
			args: args{hostname: "192.168.1.1"},
			want: "192.168.1.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetHostByName(tt.args.hostname); got != tt.want {
				t.Errorf("GetHostByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
