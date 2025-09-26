package mock

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestGet_11(t *testing.T) {
	ctx := context.Background()
	store := &SessionStore{
		sid:    "test",
		values: make(map[interface{}]interface{}),
	}

	store.values["key1"] = "value1"
	store.values[123] = "value2"

	tests := []struct {
		name string
		key  interface{}
		want interface{}
	}{
		{"string key", "key1", "value1"},
		{"int key", 123, "value2"},
		{"not exist key", "key3", nil},
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
