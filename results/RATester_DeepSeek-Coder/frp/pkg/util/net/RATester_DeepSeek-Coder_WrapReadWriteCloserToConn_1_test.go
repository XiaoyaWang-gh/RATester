package net

import (
	"fmt"
	"io"
	"net"
	"reflect"
	"testing"
)

func TestWrapReadWriteCloserToConn_1(t *testing.T) {
	type args struct {
		rwc       io.ReadWriteCloser
		underConn net.Conn
	}
	tests := []struct {
		name string
		args args
		want *WrapReadWriteCloserConn
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

			if got := WrapReadWriteCloserToConn(tt.args.rwc, tt.args.underConn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapReadWriteCloserToConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
