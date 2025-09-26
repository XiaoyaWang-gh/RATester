package gin

import (
	"fmt"
	"testing"
)

func TestFullPath_1(t *testing.T) {
	tests := []struct {
		name string
		ctx  *Context
		want string
	}{
		{
			name: "Test case 1",
			ctx: &Context{
				fullPath: "/test",
			},
			want: "/test",
		},
		{
			name: "Test case 2",
			ctx: &Context{
				fullPath: "/test2",
			},
			want: "/test2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.ctx.FullPath(); got != tt.want {
				t.Errorf("FullPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
