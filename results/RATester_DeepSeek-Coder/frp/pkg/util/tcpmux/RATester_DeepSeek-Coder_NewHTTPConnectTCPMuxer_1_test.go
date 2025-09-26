package tcpmux

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestNewHTTPConnectTCPMuxer_1(t *testing.T) {
	type args struct {
		listener    net.Listener
		passthrough bool
		timeout     time.Duration
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

			got, err := NewHTTPConnectTCPMuxer(tt.args.listener, tt.args.passthrough, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHTTPConnectTCPMuxer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && !tt.wantErr {
				t.Errorf("NewHTTPConnectTCPMuxer() = %v, want non-nil", got)
			}
		})
	}
}
