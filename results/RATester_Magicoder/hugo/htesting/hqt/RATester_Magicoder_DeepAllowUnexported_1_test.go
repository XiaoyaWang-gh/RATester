package hqt

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDeepAllowUnexported_1(t *testing.T) {
	type TestStruct struct {
		Exported   string
		unexported int
	}

	testCases := []struct {
		name     string
		input    []any
		expected cmp.Option
	}{
		{
			name: "Test with exported and unexported fields",
			input: []any{
				TestStruct{
					Exported:   "exported",
					unexported: 1,
				},
			},
			expected: cmp.AllowUnexported(TestStruct{}),
		},
		{
			name: "Test with multiple structs",
			input: []any{
				TestStruct{
					Exported:   "exported1",
					unexported: 1,
				},
				TestStruct{
					Exported:   "exported2",
					unexported: 2,
				},
			},
			expected: cmp.AllowUnexported(TestStruct{}),
		},
		{
			name: "Test with no structs",
			input: []any{
				"string",
				1,
			},
			expected: cmp.AllowUnexported(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := DeepAllowUnexported(tc.input...)
			if diff := cmp.Diff(tc.expected, result); diff != "" {
				t.Errorf("DeepAllowUnexported() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
