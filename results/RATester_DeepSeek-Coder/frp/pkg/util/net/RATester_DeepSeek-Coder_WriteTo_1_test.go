package net

import (
	"fmt"
	"net"
	"testing"
)

func TestWriteTo_1(t *testing.T) {
	type args struct {
		b    []byte
		addr net.Addr
	}
	tests := []struct {
		name    string
		c       *ConnectedUDPConn
		args    args
		wantN   int
		wantErr bool
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

			gotN, err := tt.c.WriteTo(tt.args.b, tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectedUDPConn.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("ConnectedUDPConn.WriteTo() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
