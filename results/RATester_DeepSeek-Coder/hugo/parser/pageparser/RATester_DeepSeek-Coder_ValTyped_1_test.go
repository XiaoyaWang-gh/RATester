package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValTyped_1(t *testing.T) {
	tests := []struct {
		name   string
		item   Item
		source []byte
		want   any
	}{
		{
			name: "Test case 1",
			item: Item{
				isString: true,
			},
			source: []byte("test"),
			want:   "test",
		},
		{
			name: "Test case 2",
			item: Item{
				isString: false,
			},
			source: []byte("true"),
			want:   true,
		},
		{
			name: "Test case 3",
			item: Item{
				isString: false,
			},
			source: []byte("123"),
			want:   123,
		},
		{
			name: "Test case 4",
			item: Item{
				isString: false,
			},
			source: []byte("123.456"),
			want:   123.456,
		},
		{
			name: "Test case 5",
			item: Item{
				isString: false,
			},
			source: []byte("false"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.item.ValTyped(tt.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValTyped() = %v, want %v", got, tt.want)
			}
		})
	}
}
