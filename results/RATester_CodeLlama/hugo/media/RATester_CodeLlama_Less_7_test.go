package media

import (
	"fmt"
	"testing"
)

func TestLess_7(t *testing.T) {
	t.Parallel()
	var tt = []struct {
		name string
		t    Types
		i    int
		j    int
		want bool
	}{
		{
			name: "less",
			t:    Types{{Type: "a"}, {Type: "b"}, {Type: "c"}},
			i:    0,
			j:    1,
			want: true,
		},
		{
			name: "greater",
			t:    Types{{Type: "a"}, {Type: "b"}, {Type: "c"}},
			i:    1,
			j:    0,
			want: false,
		},
		{
			name: "equal",
			t:    Types{{Type: "a"}, {Type: "b"}, {Type: "c"}},
			i:    0,
			j:    0,
			want: false,
		},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()
			if got := tc.t.Less(tc.i, tc.j); got != tc.want {
				t.Errorf("Less() = %v, want %v", got, tc.want)
			}
		})
	}
}
