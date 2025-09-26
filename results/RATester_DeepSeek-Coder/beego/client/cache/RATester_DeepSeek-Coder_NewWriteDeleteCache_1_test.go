package cache

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestNewWriteDeleteCache_1(t *testing.T) {
	type args struct {
		cache     Cache
		storeFunc func(ctx context.Context, key string, val any) error
	}
	tests := []struct {
		name    string
		args    args
		want    *WriteDeleteCache
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

			got, err := NewWriteDeleteCache(tt.args.cache, tt.args.storeFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWriteDeleteCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWriteDeleteCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
