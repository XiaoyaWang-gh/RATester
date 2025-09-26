package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestSet_28(t *testing.T) {
	type fields struct {
		Cache     Cache
		interval  time.Duration
		timeout   time.Duration
		storeFunc func(ctx context.Context, key string, val any) error
	}
	type args struct {
		ctx context.Context
		key string
		val any
	}
	tests := []struct {
		name    string
		fields  fields
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

			c := &WriteDoubleDeleteCache{
				Cache:     tt.fields.Cache,
				interval:  tt.fields.interval,
				timeout:   tt.fields.timeout,
				storeFunc: tt.fields.storeFunc,
			}
			if err := c.Set(tt.args.ctx, tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("WriteDoubleDeleteCache.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
