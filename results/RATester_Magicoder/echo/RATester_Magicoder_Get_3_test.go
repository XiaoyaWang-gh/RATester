package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_3(t *testing.T) {
	ctx := &context{
		store: Map{
			"key1": "value1",
			"key2": "value2",
		},
	}

	tests := []struct {
		name string
		key  string
		want interface{}
	}{
		{
			name: "Existing key",
			key:  "key1",
			want: "value1",
		},
		{
			name: "Non-existing key",
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

			if got := ctx.Get(tt.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
