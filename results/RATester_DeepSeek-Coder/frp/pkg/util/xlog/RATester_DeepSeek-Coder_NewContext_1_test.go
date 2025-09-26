package xlog

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestNewContext_1(t *testing.T) {
	type args struct {
		ctx context.Context
		xl  *Logger
	}
	tests := []struct {
		name string
		args args
		want context.Context
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

			if got := NewContext(tt.args.ctx, tt.args.xl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
