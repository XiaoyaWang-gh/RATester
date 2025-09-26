package tool

import (
	"fmt"
	"testing"
)

func TestTestServerPort_1(t *testing.T) {
	type args struct {
		p int
		m string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				p: 8080,
				m: "p2p",
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				p: 65536,
				m: "secret",
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				p: 80,
				m: "udp",
			},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{
				p: 443,
				m: "tcp",
			},
			want: true,
		},
		{
			name: "Test case 5",
			args: args{
				p: 1024,
				m: "p2p",
			},
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

			if got := TestServerPort(tt.args.p, tt.args.m); got != tt.want {
				t.Errorf("TestServerPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
