package media

import (
	"fmt"
	"testing"
)

func TestLess_7(t *testing.T) {
	tests := []struct {
		name string
		t    Types
		i    int
		j    int
		want bool
	}{
		{
			name: "Test case 1",
			t:    Types{{Type: "application/json"}, {Type: "text/plain"}},
			i:    0,
			j:    1,
			want: true,
		},
		{
			name: "Test case 2",
			t:    Types{{Type: "text/plain"}, {Type: "application/json"}},
			i:    0,
			j:    1,
			want: false,
		},
		{
			name: "Test case 3",
			t:    Types{{Type: "application/json"}, {Type: "application/json"}},
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

			if got := tt.t.Less(tt.i, tt.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}
