package crypt

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestNewTlsServerConn_1(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name string
		args args
		want net.Conn
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

			if got := NewTlsServerConn(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTlsServerConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
