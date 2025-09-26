package healthcheck

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestCheckHealthHTTP_1(t *testing.T) {
	type args struct {
		ctx    context.Context
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

			shc := &ServiceHealthChecker{
				client: &http.Client{},
			}
			if err := shc.checkHealthHTTP(tt.args.ctx, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("ServiceHealthChecker.checkHealthHTTP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
