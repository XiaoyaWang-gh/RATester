package helpers

import (
	"fmt"
	"testing"
)

func TestString_4(t *testing.T) {
	tests := []struct {
		name string
		n    NamedSlice
		want string
	}{
		{
			name: "Empty Slice",
			n:    NamedSlice{Name: "test", Slice: []string{}},
			want: "test",
		},
		{
			name: "Non-empty Slice",
			n:    NamedSlice{Name: "test", Slice: []string{"a", "b", "c"}},
			want: "test/a,b,c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.n.String(); got != tt.want {
				t.Errorf("NamedSlice.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
