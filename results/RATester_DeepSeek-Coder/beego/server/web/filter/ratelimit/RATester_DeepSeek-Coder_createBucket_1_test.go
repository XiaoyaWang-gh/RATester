package ratelimit

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateBucket_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		l    *limiter
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

			if got := tt.l.createBucket(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("limiter.createBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}
