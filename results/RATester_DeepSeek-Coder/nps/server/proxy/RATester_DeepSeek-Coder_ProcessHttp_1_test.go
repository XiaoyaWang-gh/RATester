package proxy

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/conn"
)

func TestProcessHttp_1(t *testing.T) {
	type args struct {
		c *conn.Conn
		s *TunnelModeServer
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

			if err := ProcessHttp(tt.args.c, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("ProcessHttp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
