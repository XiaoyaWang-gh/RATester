package redis_sentinel

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_10(t *testing.T) {
	type args struct {
		ctx         context.Context
		maxlifetime int64
		cfgStr      string
	}
	tests := []struct {
		name    string
		rp      *Provider
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

			rp := &Provider{}
			if err := rp.SessionInit(tt.args.ctx, tt.args.maxlifetime, tt.args.cfgStr); (err != nil) != tt.wantErr {
				t.Errorf("Provider.SessionInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
