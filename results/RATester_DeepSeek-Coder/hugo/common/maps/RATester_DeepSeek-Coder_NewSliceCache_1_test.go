package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewSliceCache_1(t *testing.T) {
	t.Parallel()

	type testType struct {
		name string
		want *SliceCache[any]
	}

	tests := []testType{
		{
			name: "TestNewSliceCache",
			want: &SliceCache[any]{m: make(map[string][]any)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewSliceCache[any]()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSliceCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
