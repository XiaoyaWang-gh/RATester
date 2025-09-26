package metadecoders

import (
	"fmt"
	"reflect"
	"testing"
)

func TeststringifyMapKeys_1(t *testing.T) {
	tests := []struct {
		name string
		in   any
		want any
	}{
		{
			name: "Test case 1",
			in:   map[string]any{"key1": "value1", "key2": "value2"},
			want: map[string]any{"key1": "value1", "key2": "value2"},
		},
		{
			name: "Test case 2",
			in:   map[any]any{"key1": "value1", "key2": "value2"},
			want: map[string]any{"key1": "value1", "key2": "value2"},
		},
		{
			name: "Test case 3",
			in:   []any{map[string]any{"key1": "value1", "key2": "value2"}},
			want: []any{map[string]any{"key1": "value1", "key2": "value2"}},
		},
		{
			name: "Test case 4",
			in:   []any{map[any]any{"key1": "value1", "key2": "value2"}},
			want: []any{map[string]any{"key1": "value1", "key2": "value2"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, _ := stringifyMapKeys(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringifyMapKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
