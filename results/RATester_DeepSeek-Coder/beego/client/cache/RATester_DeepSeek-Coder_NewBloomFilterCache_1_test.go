package cache

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewBloomFilterCache_1(t *testing.T) {
	type args struct {
		cache      Cache
		ln         func(ctx context.Context, key string) (any, error)
		blm        BloomFilter
		expiration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    *BloomFilterCache
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

			got, err := NewBloomFilterCache(tt.args.cache, tt.args.ln, tt.args.blm, tt.args.expiration)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBloomFilterCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBloomFilterCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
