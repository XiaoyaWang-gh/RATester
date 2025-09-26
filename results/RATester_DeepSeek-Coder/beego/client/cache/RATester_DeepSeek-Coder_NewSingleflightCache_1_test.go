package cache

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewSingleflightCache_1(t *testing.T) {
	type args struct {
		c          Cache
		expiration time.Duration
		loadFunc   func(ctx context.Context, key string) (any, error)
	}
	tests := []struct {
		name    string
		args    args
		want    Cache
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

			got, err := NewSingleflightCache(tt.args.c, tt.args.expiration, tt.args.loadFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSingleflightCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSingleflightCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
