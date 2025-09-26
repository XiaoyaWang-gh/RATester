package mock

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestmockFromCtx_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want []*Mock
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

			if got := mockFromCtx(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mockFromCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
