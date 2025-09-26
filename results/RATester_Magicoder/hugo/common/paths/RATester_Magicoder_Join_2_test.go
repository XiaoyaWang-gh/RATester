package paths

import (
	"fmt"
	"testing"
)

func TestJoin_2(t *testing.T) {
	bridge := pathBridge{}
	tests := []struct {
		name string
		in   []string
		want string
	}{
		{
			name: "simple join",
			in:   []string{"a", "b", "c"},
			want: "a/b/c",
		},
		{
			name: "empty join",
			in:   []string{},
			want: "",
		},
		{
			name: "single element join",
			in:   []string{"a"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := bridge.Join(tt.in...); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}
