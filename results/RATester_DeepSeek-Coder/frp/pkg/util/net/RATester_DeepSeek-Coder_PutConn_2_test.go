package net

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestPutConn_2(t *testing.T) {
	type fields struct {
		acceptCh chan net.Conn
		closed   bool
		mu       sync.Mutex
	}
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
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

			l := &InternalListener{
				acceptCh: tt.fields.acceptCh,
				closed:   tt.fields.closed,
				mu:       tt.fields.mu,
			}
			if err := l.PutConn(tt.args.conn); (err != nil) != tt.wantErr {
				t.Errorf("PutConn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
