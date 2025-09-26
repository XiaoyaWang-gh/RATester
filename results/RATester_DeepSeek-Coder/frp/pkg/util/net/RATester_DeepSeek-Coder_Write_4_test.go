package net

import (
	"fmt"
	"net"
	"sync"
	"testing"
	"time"
)

func TestWrite_4(t *testing.T) {
	type fields struct {
		l          *UDPListener
		localAddr  net.Addr
		remoteAddr net.Addr
		packets    chan []byte
		closeFlag  bool
		lastActive time.Time
		mu         sync.RWMutex
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
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

			c := &FakeUDPConn{
				l:          tt.fields.l,
				localAddr:  tt.fields.localAddr,
				remoteAddr: tt.fields.remoteAddr,
				packets:    tt.fields.packets,
				closeFlag:  tt.fields.closeFlag,
				lastActive: tt.fields.lastActive,
				mu:         tt.fields.mu,
			}
			gotN, err := c.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("FakeUDPConn.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("FakeUDPConn.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
