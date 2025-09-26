package ssdb

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestGet_41(t *testing.T) {
	type args struct {
		ctx context.Context
		key interface{}
	}
	tests := []struct {
		name string
		s    *SessionStore
		args args
		want interface{}
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

			if got := tt.s.Get(tt.args.ctx, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SessionStore.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
