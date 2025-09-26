package ssdb

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_12(t *testing.T) {
	type args struct {
		ctx         context.Context
		maxLifetime int64
		cfg         string
	}
	tests := []struct {
		name    string
		p       *Provider
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

			if err := tt.p.SessionInit(tt.args.ctx, tt.args.maxLifetime, tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Provider.SessionInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
