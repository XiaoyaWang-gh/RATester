package ratelimit

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestDefaultSessionKey_1(t *testing.T) {
	type args struct {
		ctx *context.Context
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := defaultSessionKey(tt.args.ctx); got != tt.want {
				t.Errorf("defaultSessionKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
