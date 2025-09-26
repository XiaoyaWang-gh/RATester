package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHandler_1(t *testing.T) {
	tests := []struct {
		name string
		c    *Context
		want HandlerFunc
	}{
		{
			name: "Test Case 1",
			c: &Context{
				handlers: HandlersChain{
					func(c *Context) {
						// some logic
					},
				},
			},
			want: func(c *Context) {
				// some logic
			},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.Handler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Context.Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}
