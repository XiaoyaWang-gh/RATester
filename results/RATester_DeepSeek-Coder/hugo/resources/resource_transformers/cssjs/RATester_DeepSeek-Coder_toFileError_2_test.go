package cssjs

import (
	"errors"
	"fmt"
	"testing"
)

func TestToFileError_2(t *testing.T) {
	type testCase struct {
		name     string
		output   string
		expected error
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			output:   "cssSyntaxErrorRe.FindStringSubmatch(output) is nil",
			expected: errors.New("cssSyntaxErrorRe.FindStringSubmatch(output) is nil"),
		},
		{
			name:     "Test case 2",
			output:   "strconv.Atoi(match[1]) returns an error",
			expected: errors.New("strconv.Atoi(match[1]) returns an error"),
		},
		{
			name:     "Test case 3",
			output:   "imp.linemap[lineNum] does not exist",
			expected: errors.New("imp.linemap[lineNum] does not exist"),
		},
		{
			name:     "Test case 4",
			output:   "imp.fs.Stat(file.Filename) returns an error",
			expected: errors.New("imp.fs.Stat(file.Filename) returns an error"),
		},
		{
			name:     "Test case 5",
			output:   "meta.Open() returns an error",
			expected: errors.New("meta.Open() returns an error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			imp := &importResolver{}
			err := imp.toFileError(tc.output)
			if err.Error() != tc.expected.Error() {
				t.Errorf("Expected error %v, but got %v", tc.expected, err)
			}
		})
	}
}
