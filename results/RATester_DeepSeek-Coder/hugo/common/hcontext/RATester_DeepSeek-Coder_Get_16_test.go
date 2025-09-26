package hcontext

import (
	"context"
	"fmt"
	"testing"
)

func TestGet_16(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		f    keyInContext[int, string]
		args args
		want int
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

			if got := tt.f.Get(tt.args.ctx); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
