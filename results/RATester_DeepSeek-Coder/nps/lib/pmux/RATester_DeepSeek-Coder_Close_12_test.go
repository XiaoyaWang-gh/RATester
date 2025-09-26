package pmux

import (
	"fmt"
	"net"
	"testing"
)

func TestClose_12(t *testing.T) {
	type fields struct {
		Conn     net.Conn
		rs       []byte
		readMore bool
		start    int
	}
	tests := []struct {
		name    string
		fields  fields
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

			pConn := &PortConn{
				Conn:     tt.fields.Conn,
				rs:       tt.fields.rs,
				readMore: tt.fields.readMore,
				start:    tt.fields.start,
			}
			if err := pConn.Close(); (err != nil) != tt.wantErr {
				t.Errorf("PortConn.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
