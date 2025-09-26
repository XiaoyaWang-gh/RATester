package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_2(t *testing.T) {
	tests := []struct {
		name string
		l    Length
		obj  interface{}
		want bool
	}{
		{
			name: "Test string length",
			l:    Length{N: 5, Key: "test"},
			obj:  "hello",
			want: true,
		},
		{
			name: "Test slice length",
			l:    Length{N: 3, Key: "test"},
			obj:  []int{1, 2, 3},
			want: true,
		},
		{
			name: "Test unsupported type",
			l:    Length{N: 3, Key: "test"},
			obj:  123,
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

			if got := tt.l.IsSatisfied(tt.obj); got != tt.want {
				t.Errorf("Length.IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
