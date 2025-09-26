package context

import (
	"fmt"
	"testing"
)

func TestRender_2(t *testing.T) {
	testCases := []struct {
		name     string
		status   StatusCode
		expected int
	}{
		{
			name:     "TestRender_200",
			status:   200,
			expected: 200,
		},
		{
			name:     "TestRender_404",
			status:   404,
			expected: 404,
		},
		{
			name:     "TestRender_500",
			status:   500,
			expected: 500,
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
				Output: &BeegoOutput{
					Status: tc.expected,
				},
			}

			tc.status.Render(ctx)

			if ctx.Output.Status != tc.expected {
				t.Errorf("Expected status code %d, got %d", tc.expected, ctx.Output.Status)
			}
		})
	}
}
