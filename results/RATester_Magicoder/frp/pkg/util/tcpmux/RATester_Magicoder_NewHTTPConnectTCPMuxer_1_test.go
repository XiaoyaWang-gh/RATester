package tcpmux

import (
	"fmt"
	"net"
	"reflect"
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
		want    *HTTPConnectTCPMuxer
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPConnectTCPMuxer() = %v, want %v", got, tt.want)
			}
		})
	}
}
