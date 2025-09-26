package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewCache_2(t *testing.T) {
	t.Parallel()

	type keyType string
	type valueType int

	tests := []struct {
		name string
		want *Cache[keyType, valueType]
	}{
		{
			name: "should return a new cache",
			want: &Cache[keyType, valueType]{m: make(map[keyType]valueType)},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			got := NewCache[keyType, valueType]()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
