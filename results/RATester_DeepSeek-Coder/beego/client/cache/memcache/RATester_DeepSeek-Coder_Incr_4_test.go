package memcache

import (
	"context"
	"fmt"
	"testing"
)

func TestIncr_4(t *testing.T) {
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		rc      *Cache
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

			if err := tt.rc.Incr(tt.args.ctx, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Cache.Incr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
