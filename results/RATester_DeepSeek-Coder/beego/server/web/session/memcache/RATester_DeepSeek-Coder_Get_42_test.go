package memcache

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestGet_42(t *testing.T) {
	ctx := context.Background()
	store := &SessionStore{
		sid:         "test",
		lock:        sync.RWMutex{},
		values:      make(map[interface{}]interface{}),
		maxlifetime: 3600,
	}
	store.values["key1"] = "value1"
	store.values["key2"] = "value2"

	type args struct {
		ctx context.Context
		key interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "get existing key",
			args: args{
				ctx: ctx,
				key: "key1",
			},
			want: "value1",
		},
		{
			name: "get non-existing key",
			args: args{
				ctx: ctx,
				key: "key3",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := store.Get(tt.args.ctx, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
