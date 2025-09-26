package client

import (
	"fmt"
	"testing"
)

func TestHandleP2PUdp_1(t *testing.T) {
	type args struct {
		localAddr   string
		rAddr       string
		md5Password string
		role        string
	}
	tests := []struct {
		name    string
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

			gotRemoteAddress, _, err := handleP2PUdp(tt.args.localAddr, tt.args.rAddr, tt.args.md5Password, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleP2PUdp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRemoteAddress == "" && !tt.wantErr {
				t.Errorf("handleP2PUdp() gotRemoteAddress = %v, want non-empty", gotRemoteAddress)
			}
		})
	}
}
