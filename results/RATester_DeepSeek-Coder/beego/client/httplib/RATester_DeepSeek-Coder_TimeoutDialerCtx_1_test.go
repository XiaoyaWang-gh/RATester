package httplib

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTimeoutDialerCtx_1(t *testing.T) {
	type args struct {
		ctx       context.Context
		netw      string
		addr      string
		cTimeout  time.Duration
		rwTimeout time.Duration
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

			got := TimeoutDialerCtx(tt.args.cTimeout, tt.args.rwTimeout)
			conn, err := got(tt.args.ctx, tt.args.netw, tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeoutDialerCtx() = %v, want %v", err, tt.wantErr)
				return
			}
			if err == nil {
				defer conn.Close()
			}
		})
	}
}
