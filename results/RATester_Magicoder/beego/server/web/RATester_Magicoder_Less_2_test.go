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
			p:    ControllerCommentsSlice{ControllerComments{Router: "a"}, ControllerComments{Router: "b"}},
			i:    0,
			j:    1,
			want: true,
		},
		{
			name: "Test case 2",
			p:    ControllerCommentsSlice{ControllerComments{Router: "b"}, ControllerComments{Router: "a"}},
			i:    0,
			j:    1,
			want: false,
		},
		{
			name: "Test case 3",
			p:    ControllerCommentsSlice{ControllerComments{Router: "a"}, ControllerComments{Router: "a"}},
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
