package nathole

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/fatedier/frp/pkg/transport"
)

func TestPreCheck_1(t *testing.T) {
	type args struct {
		ctx         context.Context
		transporter transport.MessageTransporter
		proxyName   string
		timeout     time.Duration
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

			if err := PreCheck(tt.args.ctx, tt.args.transporter, tt.args.proxyName, tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("PreCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
