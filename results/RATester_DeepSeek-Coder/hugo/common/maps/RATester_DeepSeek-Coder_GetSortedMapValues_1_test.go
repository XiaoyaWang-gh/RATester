package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetSortedMapValues_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		key      string
		mapValue map[string]any
		want     []any
	}{
		{
			name:     "Returns sorted values of a map",
			key:      "testKey",
			mapValue: map[string]any{"b": 2, "a": 1, "c": 3},
			want:     []any{1, 2, 3},
		},
		{
			name:     "Returns nil if key does not exist",
			key:      "nonExistentKey",
			mapValue: map[string]any{"b": 2, "a": 1, "c": 3},
			want:     nil,
		},
		{
			name:     "Returns nil if value is not a map",
			key:      "testKey",
			mapValue: map[string]any{"b": 2, "a": 1, "c": "3"},
			want:     nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			s := &Scratch{
				values: map[string]any{tc.key: tc.mapValue},
			}

			got := s.GetSortedMapValues(tc.key)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}
