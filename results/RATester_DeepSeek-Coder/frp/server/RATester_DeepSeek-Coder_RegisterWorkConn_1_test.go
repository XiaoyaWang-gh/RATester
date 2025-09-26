package server

import (
	"fmt"
	"net"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestRegisterWorkConn_1(t *testing.T) {
	type args struct {
		workConn net.Conn
		newMsg   *msg.NewWorkConn
	}
	tests := []struct {
		name    string
		svr     *Service
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

			if err := tt.svr.RegisterWorkConn(tt.args.workConn, tt.args.newMsg); (err != nil) != tt.wantErr {
				t.Errorf("Service.RegisterWorkConn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
