package ledis

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_3(t *testing.T) {
	type args struct {
		ctx         context.Context
		maxlifetime int64
		cfgStr      string
	}
	tests := []struct {
		name    string
		lp      *Provider
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

			if err := tt.lp.SessionInit(tt.args.ctx, tt.args.maxlifetime, tt.args.cfgStr); (err != nil) != tt.wantErr {
				t.Errorf("Provider.SessionInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
