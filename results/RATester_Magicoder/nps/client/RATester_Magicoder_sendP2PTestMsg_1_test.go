package client

import (
	"fmt"
	"net"
	"testing"
)

func TestsendP2PTestMsg_1(t *testing.T) {
	type args struct {
		localConn   *net.UDPConn
		remoteAddr1 string
		remoteAddr2 string
		remoteAddr3 string
	}
	tests := []struct {
		name    string
		args    args
		want    string
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

			got, err := sendP2PTestMsg(tt.args.localConn, tt.args.remoteAddr1, tt.args.remoteAddr2, tt.args.remoteAddr3)
			if (err != nil) != tt.wantErr {
				t.Errorf("sendP2PTestMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sendP2PTestMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
