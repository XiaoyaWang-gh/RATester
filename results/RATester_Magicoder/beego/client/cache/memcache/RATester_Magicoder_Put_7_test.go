package memcache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestPut_7(t *testing.T) {
	type args struct {
		ctx     context.Context
		key     string
		val     interface{}
		timeout time.Duration
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

			if err := tt.rc.Put(tt.args.ctx, tt.args.key, tt.args.val, tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("Cache.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
