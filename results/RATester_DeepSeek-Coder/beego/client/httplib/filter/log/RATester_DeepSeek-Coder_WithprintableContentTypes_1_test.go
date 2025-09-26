package log

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestWithprintableContentTypes_1(t *testing.T) {
	testCases := []struct {
		name   string
		input  []string
		expect []string
	}{
		{
			name:   "Test case 1",
			input:  []string{"text/plain", "text/html"},
			expect: []string{"text/plain", "text/html"},
		},
		{
			name:   "Test case 2",
			input:  []string{"application/json"},
			expect: []string{"application/json"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			builder := &FilterChainBuilder{}
			option := WithprintableContentTypes(tc.input)
			option(builder)
			assert.Equal(t, tc.expect, builder.printableContentTypes)
		})
	}
}
