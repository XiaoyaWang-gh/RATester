package net

import (
	"crypto/tls"
	"fmt"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestCheckAndEnableTLSServerConnWithTimeout_1(t *testing.T) {
	type args struct {
		c         net.Conn
		tlsConfig *tls.Config
		tlsOnly   bool
		timeout   time.Duration
	}
	tests := []struct {
		name       string
		args       args
		wantOut    net.Conn
		wantIsTLS  bool
		wantCustom bool
		wantErr    bool
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

			gotOut, gotIsTLS, gotCustom, err := CheckAndEnableTLSServerConnWithTimeout(tt.args.c, tt.args.tlsConfig, tt.args.tlsOnly, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckAndEnableTLSServerConnWithTimeout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("CheckAndEnableTLSServerConnWithTimeout() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotIsTLS != tt.wantIsTLS {
				t.Errorf("CheckAndEnableTLSServerConnWithTimeout() gotIsTLS = %v, want %v", gotIsTLS, tt.wantIsTLS)
			}
			if gotCustom != tt.wantCustom {
				t.Errorf("CheckAndEnableTLSServerConnWithTimeout() gotCustom = %v, want %v", gotCustom, tt.wantCustom)
			}
		})
	}
}
