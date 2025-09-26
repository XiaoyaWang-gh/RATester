package server

import (
	"fmt"
	"net"
	"testing"
)

func TestGetConnKey_1(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := getConnKey(tt.args.conn); got != tt.want {
				t.Errorf("getConnKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
