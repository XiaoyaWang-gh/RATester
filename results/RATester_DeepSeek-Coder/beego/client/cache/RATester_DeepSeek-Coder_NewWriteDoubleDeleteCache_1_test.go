package cache

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewWriteDoubleDeleteCache_1(t *testing.T) {
	type args struct {
		cache     Cache
		interval  time.Duration
		timeout   time.Duration
		storeFunc func(ctx context.Context, key string, val any) error
	}
	tests := []struct {
		name    string
		args    args
		want    *WriteDoubleDeleteCache
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

			got, err := NewWriteDoubleDeleteCache(tt.args.cache, tt.args.interval, tt.args.timeout, tt.args.storeFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWriteDoubleDeleteCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWriteDoubleDeleteCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
