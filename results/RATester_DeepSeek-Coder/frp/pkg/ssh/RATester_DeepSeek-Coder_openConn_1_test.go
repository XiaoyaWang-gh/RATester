package ssh

import (
	"fmt"
	"testing"
)

func TestOpenConn_1(t *testing.T) {
	type args struct {
		addr *tcpipForward
	}
	tests := []struct {
		name    string
		s       *TunnelServer
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

			got, err := tt.s.openConn(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("TunnelServer.openConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("TunnelServer.openConn() = %v, want non-nil", got)
			}
		})
	}
}
