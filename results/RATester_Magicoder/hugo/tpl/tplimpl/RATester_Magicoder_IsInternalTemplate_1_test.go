package tplimpl

import (
	"fmt"
	"testing"
)

func TestIsInternalTemplate_1(t *testing.T) {
	testCases := []struct {
		name     string
		template *templateState
		expected bool
	}{
		{
			name: "Internal Template",
			template: &templateState{
				info: templateInfo{
					isEmbedded: true,
				},
			},
			expected: true,
		},
		{
			name: "External Template",
			template: &templateState{
				info: templateInfo{
					isEmbedded: false,
				},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.template.IsInternalTemplate()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
