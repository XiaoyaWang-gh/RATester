package web

import (
	"fmt"
	"testing"
)

func TestLess_2(t *testing.T) {
	tests := []struct {
		name string
		p    ControllerCommentsSlice
		i    int
		j    int
		want bool
	}{
		{
			name: "Test case 1",
			p:    ControllerCommentsSlice{{Router: "router1"}, {Router: "router2"}},
			i:    0,
			j:    1,
			want: true,
		},
		{
			name: "Test case 2",
			p:    ControllerCommentsSlice{{Router: "router1"}, {Router: "router2"}},
			i:    1,
			j:    0,
			want: false,
		},
		{
			name: "Test case 3",
			p:    ControllerCommentsSlice{{Router: "router1"}, {Router: "router1"}},
			i:    0,
			j:    1,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.Less(tt.i, tt.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}
