package predicate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAnd_1(t *testing.T) {
	type testCase[T any] struct {
		name string
		p    P[T]
		ps   []P[T]
		want P[T]
	}

	tests := []testCase[int]{
		{
			name: "And with nil",
			p:    nil,
			ps:   []P[int]{},
			want: func(v int) bool { return true },
		},
		{
			name: "And with one predicate",
			p:    func(v int) bool { return v > 0 },
			ps:   []P[int]{},
			want: func(v int) bool { return v > 0 },
		},
		{
			name: "And with multiple predicates",
			p:    func(v int) bool { return v > 0 },
			ps:   []P[int]{func(v int) bool { return v < 10 }},
			want: func(v int) bool { return v > 0 && v < 10 },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.p.And(tt.ps...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
