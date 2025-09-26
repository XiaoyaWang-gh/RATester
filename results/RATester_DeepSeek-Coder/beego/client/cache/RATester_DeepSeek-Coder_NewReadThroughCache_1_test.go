package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewReadThroughCache_1(t *testing.T) {
	type args struct {
		cache      Cache
		expiration time.Duration
		loadFunc   func(ctx context.Context, key string) (any, error)
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

			got, err := NewReadThroughCache(tt.args.cache, tt.args.expiration, tt.args.loadFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewReadThroughCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("NewReadThroughCache() = %v, want non-nil", got)
			}
		})
	}
}
