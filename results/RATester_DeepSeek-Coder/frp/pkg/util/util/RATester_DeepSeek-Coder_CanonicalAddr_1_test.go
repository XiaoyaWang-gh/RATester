package util

import (
	"fmt"
	"testing"
)

func TestCanonicalAddr_1(t *testing.T) {
	type test struct {
		host string
		port int
		want string
	}

	tests := []test{
		{"localhost", 80, "localhost"},
		{"localhost", 443, "localhost"},
		{"localhost", 8080, "localhost:8080"},
		{"127.0.0.1", 80, "127.0.0.1"},
		{"127.0.0.1", 443, "127.0.0.1"},
		{"127.0.0.1", 8080, "127.0.0.1:8080"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s:%d", tt.host, tt.port), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := CanonicalAddr(tt.host, tt.port)
			if got != tt.want {
				t.Errorf("CanonicalAddr(%s, %d) = %s; want %s", tt.host, tt.port, got, tt.want)
			}
		})
	}
}
