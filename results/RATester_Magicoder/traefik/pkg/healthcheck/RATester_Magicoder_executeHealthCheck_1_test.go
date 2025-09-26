package healthcheck

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestExecuteHealthCheck_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		config *dynamic.ServerHealthCheck
		target *url.URL
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
			if err := shc.executeHealthCheck(tt.args.ctx, tt.args.config, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("ServiceHealthChecker.executeHealthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
