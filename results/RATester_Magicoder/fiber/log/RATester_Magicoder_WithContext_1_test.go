package log

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestWithContext_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want CommonLogger
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

			if got := WithContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
