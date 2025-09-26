package gin

import (
	"fmt"
	"testing"
)

func TestIsAborted_1(t *testing.T) {
	tests := []struct {
		name string
		c    *Context
		want bool
	}{
		{
			name: "Test case 1",
			c: &Context{
				index: 0,
			},
			want: false,
		},
		{
			name: "Test case 2",
			c: &Context{
				index: 10,
			},
			want: true,
		},
		{
			name: "Test case 3",
			c: &Context{
				index: -1,
			},
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

			if got := tt.c.IsAborted(); got != tt.want {
				t.Errorf("IsAborted() = %v, want %v", got, tt.want)
			}
		})
	}
}
