package healthcheck

import (
	"context"
	"fmt"
	"net/url"
	"testing"
)

func TestCheckHealthGRPC_1(t *testing.T) {
	type args struct {
		ctx       context.Context
		serverURL *url.URL
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

			shc := &ServiceHealthChecker{}
			if err := shc.checkHealthGRPC(tt.args.ctx, tt.args.serverURL); (err != nil) != tt.wantErr {
				t.Errorf("ServiceHealthChecker.checkHealthGRPC() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
