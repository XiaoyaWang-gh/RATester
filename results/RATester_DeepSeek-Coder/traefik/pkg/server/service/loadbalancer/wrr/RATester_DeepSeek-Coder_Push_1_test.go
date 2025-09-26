package wrr

import (
	"fmt"
	"testing"
)

func TestPush_1(t *testing.T) {
	type testCase struct {
		name string
		b    *Balancer
		x    interface{}
		want int
	}

	testCases := []testCase{
		{
			name: "Push namedHandler",
			b:    &Balancer{},
			x:    &namedHandler{name: "test", weight: 1.0},
			want: 1,
		},
		{
			name: "Push non-namedHandler",
			b:    &Balancer{},
			x:    "not a namedHandler",
			want: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tc.b.Push(tc.x)
			got := tc.b.Len()
			if got != tc.want {
				t.Errorf("Expected length %d, got %d", tc.want, got)
			}
		})
	}
}
