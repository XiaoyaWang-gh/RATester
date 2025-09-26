package couchbase

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestGet_3(t *testing.T) {
	type args struct {
		ctx  context.Context
		key  interface{}
		want interface{}
	}

	ctx := context.Background()
	cs := &SessionStore{
		values: make(map[interface{}]interface{}),
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestGet_ExistingKey",
			args: args{
				ctx:  ctx,
				key:  "existingKey",
				want: "existingValue",
			},
		},
		{
			name: "TestGet_NonExistingKey",
			args: args{
				ctx:  ctx,
				key:  "nonExistingKey",
				want: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			cs.values[tt.args.key] = tt.args.want
			got := cs.Get(tt.args.ctx, tt.args.key)
			if !reflect.DeepEqual(got, tt.args.want) {
				t.Errorf("SessionStore.Get() = %v, want %v", got, tt.args.want)
			}
		})
	}
}
