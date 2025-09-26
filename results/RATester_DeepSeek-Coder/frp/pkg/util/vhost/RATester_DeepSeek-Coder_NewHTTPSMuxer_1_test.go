package vhost

import (
	"fmt"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestNewHTTPSMuxer_1(t *testing.T) {
	type args struct {
		listener net.Listener
		timeout  time.Duration
	}

	tests := []struct {
		name    string
		args    args
		want    *HTTPSMuxer
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

			got, err := NewHTTPSMuxer(tt.args.listener, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHTTPSMuxer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPSMuxer() = %v, want %v", got, tt.want)
			}
		})
	}
}
