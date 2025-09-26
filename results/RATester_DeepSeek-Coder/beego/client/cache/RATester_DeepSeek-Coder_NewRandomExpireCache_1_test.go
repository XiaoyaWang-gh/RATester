package cache

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewRandomExpireCache_1(t *testing.T) {
	type args struct {
		adapter Cache
		opts    []RandomExpireCacheOption
	}
	tests := []struct {
		name string
		args args
		want Cache
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

			if got := NewRandomExpireCache(tt.args.adapter, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRandomExpireCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
