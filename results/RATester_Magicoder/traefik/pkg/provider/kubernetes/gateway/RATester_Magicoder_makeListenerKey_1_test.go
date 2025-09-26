package gateway

import (
	"fmt"
	"testing"

	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestMakeListenerKey_1(t *testing.T) {
	type args struct {
		l gatev1.Listener
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				l: gatev1.Listener{
					Protocol: "HTTP",
					Hostname: func() *gatev1.Hostname {
						h := gatev1.Hostname("example.com")
						return &h
					}(),
					Port: 80,
				},
			},
			want: "HTTP|example.com|80",
		},
		{
			name: "Test case 2",
			args: args{
				l: gatev1.Listener{
					Protocol: "HTTPS",
					Hostname: func() *gatev1.Hostname {
						h := gatev1.Hostname("example.com")
						return &h
					}(),
					Port: 443,
				},
			},
			want: "HTTPS|example.com|443",
		},
		{
			name: "Test case 3",
			args: args{
				l: gatev1.Listener{
					Protocol: "TCP",
					Hostname: nil,
					Port:     8080,
				},
			},
			want: "TCP||8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := makeListenerKey(tt.args.l); got != tt.want {
				t.Errorf("makeListenerKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
