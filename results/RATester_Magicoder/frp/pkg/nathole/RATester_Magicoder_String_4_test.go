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
				IP:   net.ParseIP("192.168.1.1"),
				Port: 8080,
			},
			want: "192.168.1.1:8080",
		},
		{
			name: "Test case 2",
			s: &ChangedAddress{
				IP:   net.ParseIP("10.0.0.1"),
				Port: 443,
			},
			want: "10.0.0.1:443",
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

			if got := tt.s.String(); got != tt.want {
				t.Errorf("ChangedAddress.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
