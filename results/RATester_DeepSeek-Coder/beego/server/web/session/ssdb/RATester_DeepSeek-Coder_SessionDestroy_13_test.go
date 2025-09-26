package ssdb

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionDestroy_13(t *testing.T) {
	type args struct {
		ctx context.Context
		sid string
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

			err := tt.p.SessionDestroy(tt.args.ctx, tt.args.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionDestroy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
