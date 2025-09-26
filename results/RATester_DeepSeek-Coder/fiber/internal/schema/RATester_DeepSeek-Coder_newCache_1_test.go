package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewCache_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want *cache
	}{
		{
			name: "Test newCache",
			want: &cache{
				m:       make(map[reflect.Type]*structInfo),
				regconv: make(map[reflect.Type]Converter),
				tag:     "schema",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := newCache()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
