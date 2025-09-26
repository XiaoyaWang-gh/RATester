package mock

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestCtxWithMock_2(t *testing.T) {
	type args struct {
		ctx  context.Context
		mock []*Mock
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

			if got := CtxWithMock(tt.args.ctx, tt.args.mock...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CtxWithMock() = %v, want %v", got, tt.want)
			}
		})
	}
}
