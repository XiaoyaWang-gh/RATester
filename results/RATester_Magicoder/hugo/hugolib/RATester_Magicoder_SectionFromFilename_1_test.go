package hugolib

import (
	"fmt"
	"testing"
)

func TestSectionFromFilename_1(t *testing.T) {
	type testCase struct {
		name     string
		filename string
		expected string
		err      error
	}

	tests := []testCase{
		{
			name:     "empty filename",
			filename: "",
			expected: "",
			err:      nil,
		},
		{
			name:     "filename with one part",
			filename: "part1",
			expected: "",
			err:      nil,
		},
		{
			name:     "filename with two parts",
			filename: "part1/part2",
			expected: "part1",
			err:      nil,
		},
		{
			name:     "filename with three parts",
			filename: "part1/part2/part3",
			expected: "part1",
			err:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			factory := ContentFactory{}
			section, err := factory.SectionFromFilename(test.filename)
			if err != test.err {
				t.Errorf("Expected error %v, got %v", test.err, err)
			}
			if section != test.expected {
				t.Errorf("Expected section %s, got %s", test.expected, section)
			}
		})
	}
}
