package mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionDestroy_2(t *testing.T) {
	type args struct {
		ctx context.Context
		sid string
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

			mp := &Provider{}
			if err := mp.SessionDestroy(tt.args.ctx, tt.args.sid); (err != nil) != tt.wantErr {
				t.Errorf("Provider.SessionDestroy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
