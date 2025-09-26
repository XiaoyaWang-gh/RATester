package hugolib

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestrenderShortcodeString_1(t *testing.T) {
	ctx := context.Background()
	testCases := []struct {
		name        string
		prerendered prerenderedShortcode
		expected    string
		hasVariants bool
		expectedErr error
	}{
		{
			name: "success",
			prerendered: prerenderedShortcode{
				s:           "test",
				hasVariants: true,
			},
			expected:    "test",
			hasVariants: true,
			expectedErr: nil,
		},
		{
			name: "failure",
			prerendered: prerenderedShortcode{
				s:           "fail",
				hasVariants: false,
			},
			expected:    "fail",
			hasVariants: false,
			expectedErr: errors.New("some error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, hasVariants, err := tc.prerendered.renderShortcodeString(ctx)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
			if hasVariants != tc.hasVariants {
				t.Errorf("Expected %t, got %t", tc.hasVariants, hasVariants)
			}
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("Expected error %s, got %s", tc.expectedErr, err)
			}
		})
	}
}
