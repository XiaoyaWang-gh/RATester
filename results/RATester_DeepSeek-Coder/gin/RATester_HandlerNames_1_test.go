package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHandlerNames_1(t *testing.T) {
	type testCase struct {
		name     string
		handlers HandlersChain
		expected []string
	}

	testCases := []testCase{
		{
			name:     "Empty HandlersChain",
			handlers: []HandlerFunc{},
			expected: []string{},
		},
		{
			name: "Non-empty HandlersChain",
			handlers: []HandlerFunc{
				func(c *Context) {
					c.Keys["key1"] = "value1"
				},
				func(c *Context) {
					c.Keys["key2"] = "value2"
				},
			},
			expected: []string{"github.com/gin-gonic/gin.(*Context).func1", "github.com/gin-gonic/gin.(*Context).func2"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := &Context{
				handlers: tc.handlers,
			}
			result := ctx.HandlerNames()
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
