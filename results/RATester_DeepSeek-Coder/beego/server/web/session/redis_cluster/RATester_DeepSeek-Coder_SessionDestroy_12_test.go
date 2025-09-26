package redis_cluster

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionDestroy_12(t *testing.T) {
	type args struct {
		ctx context.Context
		sid string
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

			err := tt.rp.SessionDestroy(tt.args.ctx, tt.args.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionDestroy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
