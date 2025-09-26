package tcpmux

import (
	"fmt"
	"net"
	"testing"
)

func TestAuth_2(t *testing.T) {
	type args struct {
		c        net.Conn
		username string
		password string
		reqInfo  map[string]string
	}
	tests := []struct {
		name    string
		muxer   *HTTPConnectTCPMuxer
		args    args
		want    bool
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

			got, err := tt.muxer.auth(tt.args.c, tt.args.username, tt.args.password, tt.args.reqInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPConnectTCPMuxer.auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTTPConnectTCPMuxer.auth() = %v, want %v", got, tt.want)
			}
		})
	}
}
