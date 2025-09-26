package server

import (
	"fmt"
	"net"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestRegisterControl_1(t *testing.T) {
	type args struct {
		ctlConn  net.Conn
		loginMsg *msg.Login
		internal bool
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

			if err := tt.svr.RegisterControl(tt.args.ctlConn, tt.args.loginMsg, tt.args.internal); (err != nil) != tt.wantErr {
				t.Errorf("Service.RegisterControl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
