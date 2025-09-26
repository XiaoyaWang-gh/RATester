package net

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestNewContextFromConn_1(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name string
		args args
		want context.Context
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

			if got := NewContextFromConn(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContextFromConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
