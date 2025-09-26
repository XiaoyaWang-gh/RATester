package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPop_2(t *testing.T) {
	type testCase[T any] struct {
		name     string
		stack    *Stack[T]
		wantItem T
		wantOk   bool
	}
	tests := []testCase[any]{
		{
			name: "pop from empty stack",
			stack: &Stack[any]{
				items: []any{},
			},
			wantItem: nil,
			wantOk:   false,
		},
		{
			name: "pop from non-empty stack",
			stack: &Stack[any]{
				items: []any{1, 2, 3},
			},
			wantItem: 3,
			wantOk:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotItem, gotOk := tt.stack.Pop()
			if !reflect.DeepEqual(gotItem, tt.wantItem) || gotOk != tt.wantOk {
				t.Errorf("Pop() = %v, %v, want %v, %v", gotItem, gotOk, tt.wantItem, tt.wantOk)
			}
		})
	}
}
