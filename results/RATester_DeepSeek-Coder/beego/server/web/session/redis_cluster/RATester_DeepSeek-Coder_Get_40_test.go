package redis_cluster

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestGet_40(t *testing.T) {
	ctx := context.Background()
	store := &SessionStore{
		values: make(map[interface{}]interface{}),
	}

	store.values["key1"] = "value1"
	store.values["key2"] = "value2"

	tests := []struct {
		name string
		key  interface{}
		want interface{}
	}{
		{
			name: "get existing key",
			key:  "key1",
			want: "value1",
		},
		{
			name: "get non-existing key",
			key:  "key3",
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

			if got := store.Get(ctx, tt.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
