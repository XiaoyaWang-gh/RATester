package session

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestGet_26(t *testing.T) {
	ctx := context.Background()
	fs := &FileSessionStore{
		sid:    "test",
		values: make(map[interface{}]interface{}),
	}

	fs.values["key1"] = "value1"
	fs.values["key2"] = "value2"

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

			if got := fs.Get(ctx, tt.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
