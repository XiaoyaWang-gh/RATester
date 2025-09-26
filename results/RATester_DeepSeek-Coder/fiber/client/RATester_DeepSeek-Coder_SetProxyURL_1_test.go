package client

import (
	"fmt"
	"testing"
)

func TestSetProxyURL_1(t *testing.T) {
	type args struct {
		proxyURL string
	}
	tests := []struct {
		name    string
		c       *Client
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

			if err := tt.c.SetProxyURL(tt.args.proxyURL); (err != nil) != tt.wantErr {
				t.Errorf("Client.SetProxyURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
