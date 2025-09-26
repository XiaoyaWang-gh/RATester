package client

import (
	"fmt"
	"net"
	"testing"
)

func TestLogin_2(t *testing.T) {
	type args struct {
		svr *Service
	}
	tests := []struct {
		name    string
		args    args
		want    net.Conn
		want1   Connector
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

			got, got1, err := tt.args.svr.login()
			if (err != nil) != tt.wantErr {
				t.Errorf("login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("login() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("login() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
