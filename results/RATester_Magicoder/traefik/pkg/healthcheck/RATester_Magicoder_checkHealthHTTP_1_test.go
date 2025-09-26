package healthcheck

import (
	"context"
	"fmt"
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
		{
			name: "test case 1",
			args: args{
				ctx:    context.Background(),
				target: &url.URL{},
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				ctx:    context.Background(),
				target: &url.URL{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			shc := &ServiceHealthChecker{}
			if err := shc.checkHealthHTTP(tt.args.ctx, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("ServiceHealthChecker.checkHealthHTTP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
