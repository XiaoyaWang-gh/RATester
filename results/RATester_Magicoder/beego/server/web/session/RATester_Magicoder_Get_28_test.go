package session

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestGet_28(t *testing.T) {
	store := &MemSessionStore{
		value: map[interface{}]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	ctx := context.Background()

	tests := []struct {
		name string
		key  interface{}
		want interface{}
	}{
		{
			name: "existing key",
			key:  "key1",
			want: "value1",
		},
		{
			name: "non-existing key",
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
