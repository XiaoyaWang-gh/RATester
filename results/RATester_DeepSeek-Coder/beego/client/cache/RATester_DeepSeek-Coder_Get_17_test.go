package cache

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestGet_17(t *testing.T) {
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		c       *readThroughCache
		args    args
		want    any
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

			got, err := tt.c.Get(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("readThroughCache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readThroughCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
