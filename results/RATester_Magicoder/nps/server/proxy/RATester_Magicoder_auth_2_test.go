package proxy

import (
	"fmt"
	"net/http"
	"testing"

	"ehang.io/nps/lib/conn"
)

func Testauth_2(t *testing.T) {
	type args struct {
		r *http.Request
		c *conn.Conn
		u string
		p string
	}
	tests := []struct {
		name    string
		s       *BaseServer
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

			if err := tt.s.auth(tt.args.r, tt.args.c, tt.args.u, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("BaseServer.auth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
