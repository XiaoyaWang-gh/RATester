package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPush_2(t *testing.T) {
	t.Parallel()

	type testCase[T any] struct {
		name  string
		stack *Stack[T]
		item  T
		want  []T
	}

	tests := []testCase[int]{
		{
			name:  "push to empty stack",
			stack: &Stack[int]{},
			item:  1,
			want:  []int{1},
		},
		{
			name:  "push to non-empty stack",
			stack: &Stack[int]{items: []int{1, 2, 3}},
			item:  4,
			want:  []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.stack.Push(tt.item)
			if !reflect.DeepEqual(tt.stack.items, tt.want) {
				t.Errorf("got %v, want %v", tt.stack.items, tt.want)
			}
		})
	}
}
