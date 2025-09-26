package visitor

import (
	"fmt"
	"net"
	"testing"
)

func TestTransferConn_1(t *testing.T) {
	type args struct {
		name string
		conn net.Conn
	}
	tests := []struct {
		name    string
		v       *visitorHelperImpl
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

			v := &visitorHelperImpl{
				connectServerFn: tt.v.connectServerFn,
				msgTransporter:  tt.v.msgTransporter,
				transferConnFn:  tt.v.transferConnFn,
				runID:           tt.v.runID,
			}
			if err := v.TransferConn(tt.args.name, tt.args.conn); (err != nil) != tt.wantErr {
				t.Errorf("visitorHelperImpl.TransferConn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
