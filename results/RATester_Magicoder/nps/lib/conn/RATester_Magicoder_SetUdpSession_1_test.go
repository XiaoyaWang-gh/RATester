package conn

import (
	"fmt"
	"testing"

	"github.com/xtaci/kcp-go"
)

func TestSetUdpSession_1(t *testing.T) {
	type args struct {
		sess *kcp.UDPSession
	}
	tests := []struct {
		name string
		args args
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

			SetUdpSession(tt.args.sess)
		})
	}
}
