package page

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_7(t *testing.T) {
	testCases := []struct {
		name   string
		pee    *permalinkExpandError
		expect string
	}{
		{
			name: "Test case 1",
			pee: &permalinkExpandError{
				pattern: "testPattern",
				err:     errors.New("test error"),
			},
			expect: `error expanding "testPattern": test error`,
		},
		{
			name: "Test case 2",
			pee: &permalinkExpandError{
				pattern: "anotherTestPattern",
				err:     errors.New("another test error"),
			},
			expect: `error expanding "anotherTestPattern": another test error`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.pee.Error()
			if result != tc.expect {
				t.Errorf("Expected %q, got %q", tc.expect, result)
			}
		})
	}
}
