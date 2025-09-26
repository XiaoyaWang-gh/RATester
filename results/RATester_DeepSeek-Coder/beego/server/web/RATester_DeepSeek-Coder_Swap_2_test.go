package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSwap_2(t *testing.T) {
	tests := []struct {
		name string
		p    ControllerCommentsSlice
		i    int
		j    int
		want ControllerCommentsSlice
	}{
		{
			name: "Test case 1",
			p:    ControllerCommentsSlice{ControllerComments{Method: "Method1"}, ControllerComments{Method: "Method2"}},
			i:    0,
			j:    1,
			want: ControllerCommentsSlice{ControllerComments{Method: "Method2"}, ControllerComments{Method: "Method1"}},
		},
		{
			name: "Test case 2",
			p:    ControllerCommentsSlice{ControllerComments{Method: "Method1"}, ControllerComments{Method: "Method2"}},
			i:    1,
			j:    0,
			want: ControllerCommentsSlice{ControllerComments{Method: "Method2"}, ControllerComments{Method: "Method1"}},
		},
		{
			name: "Test case 3",
			p:    ControllerCommentsSlice{ControllerComments{Method: "Method1"}},
			i:    0,
			j:    0,
			want: ControllerCommentsSlice{ControllerComments{Method: "Method1"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.p.Swap(tt.i, tt.j)
			if !reflect.DeepEqual(tt.p, tt.want) {
				t.Errorf("got %v, want %v", tt.p, tt.want)
			}
		})
	}
}
