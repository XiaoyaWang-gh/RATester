package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestSet_30(t *testing.T) {
	type args struct {
		ctx context.Context
		key string
		val any
	}
	tests := []struct {
		name    string
		w       *WriteDeleteCache
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

			w := &WriteDeleteCache{
				Cache:     tt.w.Cache,
				storeFunc: tt.w.storeFunc,
			}
			if err := w.Set(tt.args.ctx, tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("WriteDeleteCache.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
