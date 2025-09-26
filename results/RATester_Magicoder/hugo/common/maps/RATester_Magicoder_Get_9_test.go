package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_9(t *testing.T) {
	scratch := &Scratch{
		values: map[string]any{
			"key1": "value1",
			"key2": "value2",
		},
	}

	tests := []struct {
		name string
		key  string
		want any
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

			if got := scratch.Get(tt.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
