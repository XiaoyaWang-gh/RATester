package client

import (
	"fmt"
	"net/url"
	"testing"
)

func TestNewHttpProxyConn_1(t *testing.T) {
	type args struct {
		url        *url.URL
		remoteAddr string
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

			got, err := NewHttpProxyConn(tt.args.url, tt.args.remoteAddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHttpProxyConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("NewHttpProxyConn() = %v, want not nil", got)
			}
		})
	}
}
