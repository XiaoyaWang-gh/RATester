package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHandler_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		ctx  *Context
		want HandlerFunc
	}{
		{
			name: "Test Case 1",
			ctx:  &Context{handlers: HandlersChain{func(c *Context) {}}},
			want: func(c *Context) {},
		},
		{
			name: "Test Case 2",
			ctx:  &Context{handlers: HandlersChain{func(c *Context) {}, func(c *Context) {}}},
			want: func(c *Context) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.ctx.Handler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Context.Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}
