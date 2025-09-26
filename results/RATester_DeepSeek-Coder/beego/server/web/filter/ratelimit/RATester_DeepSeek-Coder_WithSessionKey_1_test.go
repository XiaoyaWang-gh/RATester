package ratelimit

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestWithSessionKey_1(t *testing.T) {
	type args struct {
		f func(ctx *context.Context) string
	}
	tests := []struct {
		name string
		args args
		want limiterOption
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

			if got := WithSessionKey(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithSessionKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
