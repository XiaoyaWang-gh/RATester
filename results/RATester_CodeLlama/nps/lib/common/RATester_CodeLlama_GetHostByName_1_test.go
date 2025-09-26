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
			name: "test case 1",
			args: args{
				hostname: "www.baidu.com",
			},
			want: "180.97.33.108",
		},
		{
			name: "test case 2",
			args: args{
				hostname: "www.google.com",
			},
			want: "216.58.194.174",
		},
		{
			name: "test case 3",
			args: args{
				hostname: "www.sina.com",
			},
			want: "183.62.198.10",
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
