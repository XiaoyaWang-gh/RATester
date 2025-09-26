package conn

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestNewConn_1(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name string
		args args
		want *Conn
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

			if got := NewConn(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
