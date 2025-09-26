package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_2(t *testing.T) {
	type testCase struct {
		name          string
		ctx           *Context
		err           error
		expectedError *Error
	}

	testCases := []testCase{
		{
			name: "NilError",
			ctx:  &Context{},
			err:  nil,
			expectedError: &Error{
				Err:  nil,
				Type: ErrorTypePrivate,
			},
		},
		{
			name: "NonNilError",
			ctx:  &Context{},
			err:  errors.New("test error"),
			expectedError: &Error{
				Err:  errors.New("test error"),
				Type: ErrorTypePrivate,
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

			result := tc.ctx.Error(tc.err)
			if result.Err != tc.expectedError.Err || result.Type != tc.expectedError.Type {
				t.Errorf("Expected %v, but got %v", tc.expectedError, result)
			}
		})
	}
}
