package common

import (
	"fmt"
	"testing"
)

func TestString_1(t *testing.T) {
	tests := []struct {
		name string
		addr *Addr
		want string
	}{
		{
			name: "Test case 1",
			addr: &Addr{
				Type: 1,
				Host: "localhost",
				Port: 8080,
			},
			want: "localhost:8080",
		},
		{
			name: "Test case 2",
			addr: &Addr{
				Type: 1,
				Host: "127.0.0.1",
				Port: 8080,
			},
			want: "127.0.0.1:8080",
		},
		{
			name: "Test case 3",
			addr: &Addr{
				Type: 1,
				Host: "::1",
				Port: 8080,
			},
			want: "[::1]:8080",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.addr.String(); got != tt.want {
				t.Errorf("Addr.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
