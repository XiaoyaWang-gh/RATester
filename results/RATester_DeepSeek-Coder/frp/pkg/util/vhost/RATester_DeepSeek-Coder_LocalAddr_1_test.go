package vhost

import (
	"fmt"
	"io"
	"net"
	"reflect"
	"testing"
)

func TestLocalAddr_1(t *testing.T) {
	type fields struct {
		reader io.Reader
	}
	tests := []struct {
		name   string
		fields fields
		want   net.Addr
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

			conn := readOnlyConn{
				reader: tt.fields.reader,
			}
			if got := conn.LocalAddr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readOnlyConn.LocalAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
