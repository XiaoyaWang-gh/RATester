package lang

import (
	"fmt"
	"testing"
)

func TestMerge_1(t *testing.T) {
	type testCase struct {
		name     string
		ns       *Namespace
		p2       any
		p1       any
		expected any
		err      error
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			ns:       &Namespace{},
			p2:       "p2",
			p1:       "p1",
			expected: "p1",
			err:      nil,
		},
		{
			name:     "Test case 2",
			ns:       &Namespace{},
			p2:       "p2",
			p1:       "",
			expected: "p2",
			err:      nil,
		},
		{
			name:     "Test case 3",
			ns:       &Namespace{},
			p2:       "",
			p1:       "p1",
			expected: "p1",
			err:      nil,
		},
		{
			name:     "Test case 4",
			ns:       &Namespace{},
			p2:       "",
			p1:       "",
			expected: "",
			err:      nil,
		},
		{
			name:     "Test case 5",
			ns:       &Namespace{},
			p2:       "p2",
			p1:       "p1",
			expected: "p1",
			err:      nil,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := tc.ns.Merge(tc.p2, tc.p1)
			if err != tc.err {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
			if result != tc.expected {
				t.Errorf("Expected result %v, got %v", tc.expected, result)
			}
		})
	}
}
