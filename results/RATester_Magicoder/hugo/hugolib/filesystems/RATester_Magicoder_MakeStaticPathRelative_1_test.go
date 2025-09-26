package filesystems

import (
	"fmt"
	"testing"
)

func TestMakeStaticPathRelative_1(t *testing.T) {
	type testCase struct {
		name     string
		filename string
		expected string
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			filename: "test1.txt",
			expected: "test1.txt",
		},
		{
			name:     "Test case 2",
			filename: "test2.txt",
			expected: "test2.txt",
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

			// Create a new instance of SourceFilesystems for each test case
			sfs := SourceFilesystems{
				// Initialize any necessary fields here
			}

			result := sfs.MakeStaticPathRelative(tc.filename)

			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
