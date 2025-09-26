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
			name: "Test with two structs",
			input: []any{
				TestStruct{
					Exported:   "test",
					unexported: 1,
				},
				TestStruct{
					Exported:   "test2",
					unexported: 2,
				},
			},
			expected: cmp.AllowUnexported(TestStruct{}),
		},
		{
			name: "Test with three structs",
			input: []any{
				TestStruct{
					Exported:   "test",
					unexported: 1,
				},
				TestStruct{
					Exported:   "test2",
					unexported: 2,
				},
				TestStruct{
					Exported:   "test3",
					unexported: 3,
				},
			},
			expected: cmp.AllowUnexported(TestStruct{}),
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
			if diff := cmp.Diff(tc.expected, result, cmp.AllowUnexported(TestStruct{})); diff != "" {
				t.Errorf("DeepAllowUnexported() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
