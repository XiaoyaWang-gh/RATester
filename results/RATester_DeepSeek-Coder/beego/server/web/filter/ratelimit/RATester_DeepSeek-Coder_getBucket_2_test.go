package ratelimit

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestGetBucket_2(t *testing.T) {
	type args struct {
		ctx *context.Context
	}
	tests := []struct {
		name string
		args args
		want bucket
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

			l := &limiter{
				RWMutex: sync.RWMutex{},
				buckets: make(map[string]bucket),
			}
			if got := l.getBucket(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("limiter.getBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}
