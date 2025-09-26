package net

import (
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/fatedier/frp/pkg/util/xlog"
)

func TestNewLogFromConn_1(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name string
		args args
		want *xlog.Logger
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

			if got := NewLogFromConn(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogFromConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
