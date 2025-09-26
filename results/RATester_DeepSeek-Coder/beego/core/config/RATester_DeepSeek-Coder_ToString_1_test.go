package config

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestToString_1(t *testing.T) {
	type testCase struct {
		input    interface{}
		expected string
	}

	testCases := []testCase{
		{input: "Hello, world", expected: "Hello, world"},
		{input: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC), expected: "Monday, January 2, 2022"},
		{input: errors.New("an error"), expected: "an error"},
		{input: struct{ String string }{String: "custom string"}, expected: "custom string"},
		{input: 123, expected: "123"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input=%v", tc.input), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ToString(tc.input)
			if result != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, result)
			}
		})
	}
}
