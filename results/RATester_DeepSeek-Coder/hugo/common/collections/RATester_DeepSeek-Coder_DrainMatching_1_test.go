package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDrainMatching_1(t *testing.T) {
	t.Parallel()

	type testCase[T any] struct {
		name      string
		stack     *Stack[T]
		predicate func(T) bool
		want      []T
	}

	tests := []testCase[int]{
		{
			name: "drain matching",
			stack: &Stack[int]{
				items: []int{1, 2, 3, 4, 5},
			},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{2, 4},
		},
		{
			name: "no matching",
			stack: &Stack[int]{
				items: []int{1, 3, 5, 7, 9},
			},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{},
		},
		{
			name: "all matching",
			stack: &Stack[int]{
				items: []int{2, 4, 6, 8, 10},
			},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{2, 4, 6, 8, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.stack.DrainMatching(tt.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DrainMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
