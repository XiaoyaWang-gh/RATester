package web

import (
	"fmt"
	"testing"
)

func TestCompare_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		a, b   interface{}
		expect bool
	}

	tests := []test{
		{"abc", "abc", true},
		{"abc", "abcd", false},
		{123, 123, true},
		{123, 456, false},
		{123.456, 123.456, true},
		{123.456, 789.102, false},
		{true, true, true},
		{true, false, false},
		{nil, nil, true},
		{"", "", true},
		{" ", " ", true},
		{" \n\t", " \n\t", true},
		{"abc", "abc ", true},
		{" abc", "abc", true},
		{" abc ", "abc", true},
		{"abc", " abc", true},
		{"abc", "abc\n", true},
		{"abc\n", "abc", true},
		{"abc", "abc\t", true},
		{"abc\t", "abc", true},
		{"abc", "abc\r", true},
		{"abc\r", "abc", true},
	}

	for _, test := range tests {
		result := Compare(test.a, test.b)
		if result != test.expect {
			t.Errorf("Compare(%v, %v) = %v; want %v", test.a, test.b, result, test.expect)
		}
	}
}
