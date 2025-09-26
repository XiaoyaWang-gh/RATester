package vhost

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestNewMuxer_1(t *testing.T) {
	type args struct {
		listener  net.Listener
		vhostFunc muxFunc
		timeout   time.Duration
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

			got, err := NewMuxer(tt.args.listener, tt.args.vhostFunc, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMuxer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("NewMuxer() = %v, want non-nil", got)
			}
		})
	}
}
