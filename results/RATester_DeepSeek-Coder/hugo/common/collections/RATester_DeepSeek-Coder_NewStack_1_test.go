package collections

import (
	"fmt"
	"testing"
)

func TestNewStack_1(t *testing.T) {
	t.Parallel()

	type testType struct {
		name string
		want *Stack[any]
	}

	tests := []testType{
		{
			name: "TestNewStack",
			want: &Stack[any]{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewStack[any]()
			if got == nil {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
