package common

import (
	"fmt"
	"testing"
)

func TestTestUdpPort_1(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestUdpPort_0",
			args: args{port: 8080},
			want: true,
		},
		{
			name: "TestUdpPort_1",
			args: args{port: 8081},
			want: true,
		},
		{
			name: "TestUdpPort_2",
			args: args{port: 8082},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := TestUdpPort(tt.args.port); got != tt.want {
				t.Errorf("TestUdpPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
