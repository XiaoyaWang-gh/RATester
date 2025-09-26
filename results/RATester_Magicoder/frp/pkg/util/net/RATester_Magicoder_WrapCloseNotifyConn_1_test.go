package net

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestWrapCloseNotifyConn_1(t *testing.T) {
	type args struct {
		c       net.Conn
		closeFn func()
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

			if got := WrapCloseNotifyConn(tt.args.c, tt.args.closeFn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapCloseNotifyConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
