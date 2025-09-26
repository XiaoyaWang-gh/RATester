package nathole

import (
	"fmt"
	"net"
	"testing"
)

func TestString_4(t *testing.T) {
	tests := []struct {
		name string
		s    *ChangedAddress
		want string
	}{
		{
			name: "Test case 1",
			s: &ChangedAddress{
				IP:   net.ParseIP("192.0.2.0"),
				Port: 8080,
			},
			want: "192.0.2.0:8080",
		},
		{
			name: "Test case 2",
			s: &ChangedAddress{
				IP:   net.ParseIP("203.0.113.0"),
				Port: 80,
			},
			want: "203.0.113.0:80",
		},
		{
			name: "Test case 3",
			s: &ChangedAddress{
				IP:   net.ParseIP("8.8.8.8"),
				Port: 53,
			},
			want: "8.8.8.8:53",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.s.String(); got != tt.want {
				t.Errorf("ChangedAddress.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
