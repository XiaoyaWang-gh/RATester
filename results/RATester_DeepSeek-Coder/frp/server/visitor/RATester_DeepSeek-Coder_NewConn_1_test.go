package visitor

import (
	"fmt"
	"net"
	"testing"
)

func TestNewConn_1(t *testing.T) {
	type args struct {
		name           string
		conn           net.Conn
		timestamp      int64
		signKey        string
		useEncryption  bool
		useCompression bool
		visitorUser    string
	}
	tests := []struct {
		name    string
		vm      *Manager
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

			err := tt.vm.NewConn(tt.args.name, tt.args.conn, tt.args.timestamp, tt.args.signKey, tt.args.useEncryption, tt.args.useCompression, tt.args.visitorUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.NewConn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
