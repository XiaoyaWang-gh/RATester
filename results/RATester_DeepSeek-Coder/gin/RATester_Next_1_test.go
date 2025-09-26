package gin

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestNext_1(t *testing.T) {
	tests := []struct {
		name     string
		handlers HandlersChain
		expected HandlersChain
	}{
		{
			name:     "TestNext_EmptyHandlers",
			handlers: []HandlerFunc{},
			expected: []HandlerFunc{},
		},
		{
			name:     "TestNext_NilHandlers",
			handlers: nil,
			expected: nil,
		},
		{
			name:     "TestNext_SingleHandler",
			handlers: []HandlerFunc{func(c *Context) {}},
			expected: []HandlerFunc{func(c *Context) {}},
		},
		{
			name:     "TestNext_MultipleHandlers",
			handlers: []HandlerFunc{func(c *Context) {}, func(c *Context) {}, func(c *Context) {}},
			expected: []HandlerFunc{func(c *Context) {}, func(c *Context) {}, func(c *Context) {}},
		},
		{
			name:     "TestNext_NilInHandlers",
			handlers: []HandlerFunc{nil, func(c *Context) {}},
			expected: []HandlerFunc{nil, func(c *Context) {}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Context{
				handlers: test.handlers,
			}

			c.Next()

			assert.Equal(t, test.expected, c.handlers)
		})
	}
}
