package ssh

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"golang.org/x/crypto/ssh"
)

func TestWaitForwardAddrAndExtraPayload_1(t *testing.T) {
	type args struct {
		channels <-chan ssh.NewChannel
		requests <-chan *ssh.Request
		timeout  time.Duration
	}
	tests := []struct {
		name    string
		s       *TunnelServer
		args    args
		want    *tcpipForward
		want1   string
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

			got, got1, err := tt.s.waitForwardAddrAndExtraPayload(tt.args.channels, tt.args.requests, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("TunnelServer.waitForwardAddrAndExtraPayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TunnelServer.waitForwardAddrAndExtraPayload() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TunnelServer.waitForwardAddrAndExtraPayload() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
