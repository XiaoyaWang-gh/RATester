package resources

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestConcat_1(t *testing.T) {
	type testCase struct {
		name        string
		targetPath  any
		r           any
		expected    resource.Resource
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Concat with valid input",
			targetPath:  "valid/path",
			r:           []resource.Resource{},
			expected:    nil,
			expectedErr: errors.New("must provide one or more Resource objects to concat"),
		},
		{
			name:        "Concat with invalid targetPath",
			targetPath:  123,
			r:           []resource.Resource{},
			expected:    nil,
			expectedErr: fmt.Errorf("expected string, received int instead"),
		},
		{
			name:        "Concat with invalid r",
			targetPath:  "valid/path",
			r:           "invalid",
			expected:    nil,
			expectedErr: fmt.Errorf("expected slice of Resource objects, received string instead"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ns := &Namespace{}
			result, err := ns.Concat(tc.targetPath, tc.r)
			if err != nil && tc.expectedErr == nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
			if err == nil && tc.expectedErr != nil {
				t.Errorf("Expected error: %v, but got no error", tc.expectedErr)
			}
			if err != nil && tc.expectedErr != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("Expected error: %v, but got: %v", tc.expectedErr, err)
			}
			if result != tc.expected {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
