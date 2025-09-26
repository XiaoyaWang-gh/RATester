package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValues_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		scratch *Scratch
		want    map[string]any
	}{
		{
			name: "Returns the values in the scratch",
			scratch: &Scratch{
				values: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
			},
			want: map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "Returns an empty map when scratch is empty",
			scratch: &Scratch{
				values: map[string]any{},
			},
			want: map[string]any{},
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

			got := tc.scratch.Values()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
