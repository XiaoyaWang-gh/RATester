package pmux

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestNewPortConn_1(t *testing.T) {
	type args struct {
		conn     net.Conn
		rs       []byte
		readMore bool
	}
	tests := []struct {
		name string
		args args
		want *PortConn
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

			got := newPortConn(tt.args.conn, tt.args.rs, tt.args.readMore)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPortConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
