package echo

import (
	"fmt"
	"net"
	"testing"
)

func TestTrust_1(t *testing.T) {
	type args struct {
		ip net.IP
	}
	tests := []struct {
		name string
		c    *ipChecker
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.trust(tt.args.ip); got != tt.want {
				t.Errorf("trust() = %v, want %v", got, tt.want)
			}
		})
	}
}
