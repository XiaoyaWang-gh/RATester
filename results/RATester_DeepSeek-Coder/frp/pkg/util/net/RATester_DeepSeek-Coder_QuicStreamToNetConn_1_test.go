package net

import (
	"fmt"
	"net"
	"reflect"
	"testing"

	quic "github.com/quic-go/quic-go"
)

func TestQuicStreamToNetConn_1(t *testing.T) {
	type args struct {
		s quic.Stream
		c quic.Connection
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

			got := QuicStreamToNetConn(tt.args.s, tt.args.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuicStreamToNetConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
