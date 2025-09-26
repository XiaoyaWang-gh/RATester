package fiber

import (
	"fmt"
	"testing"
)

func TestNewError_1(t *testing.T) {
	t.Run("TestNewError", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name     string
			code     int
			message  string
			expected *Error
		}{
			{
				name:    "Test case 1",
				code:    400,
				message: "Bad Request",
				expected: &Error{
					Message: "Bad Request",
					Code:    400,
				},
			},
			{
				name:    "Test case 2",
				code:    404,
				message: "Not Found",
				expected: &Error{
					Message: "Not Found",
					Code:    404,
				},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				result := NewError(tc.code, tc.message)
				if result.Message != tc.expected.Message {
					t.Errorf("Expected message '%s', but got '%s'", tc.expected.Message, result.Message)
				}
				if result.Code != tc.expected.Code {
					t.Errorf("Expected code '%d', but got '%d'", tc.expected.Code, result.Code)
				}
			})
		}
	})
}
