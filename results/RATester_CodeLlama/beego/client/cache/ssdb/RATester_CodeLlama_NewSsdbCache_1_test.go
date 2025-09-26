package ssdb

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/cache"
)

func TestNewSsdbCache_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want cache.Cache
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

			if got := NewSsdbCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSsdbCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
